package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/pkoukk/chatgpt-cli/gpt"
	"github.com/pkoukk/survey/v2"
)

type CLI struct {
	accounts    []*gpt.SecretKeyAccount
	config      *gpt.GlobalConfig
	conversions []*gpt.Conversion

	selectedAccount    *gpt.SecretKeyAccount
	selectedConversion *gpt.Conversion
	ac                 *gpt.ApiConversion
}

func NewCLI() *CLI {
	cli := &CLI{}
	cli.init()
	return cli
}

func (u *CLI) SelectAccount() {
start:
	options := []any{}
	for _, account := range u.accounts {
		options = append(options, account)
	}
	addNewOption := getLanguageDes(Key_OPTION_AddNewAccount)
	options = append(options, addNewOption)
	pro := promptui.Select{
		Label: getLanguageDes(Key_LABEL_SelectAccount),
		Items: options,
	}
	id, result, err := pro.Run()
	if err != nil {
		log.Fatalln("Failed to select account:", err)
	}
	if result == addNewOption {
		q := []*survey.Question{
			{
				Name: "AccountName",
				Prompt: &survey.Input{
					Message: getLanguageDes(Key_MESSAGE_AccountName),
				},
				Validate: survey.Required,
			},
			{
				Name: "SecretKey",
				Prompt: &survey.Input{
					Message: getLanguageDes(Key_MESSAGE_SecretKey),
				},
				Validate: survey.Required,
			},
		}
		var answers gpt.SecretKeyAccount
		err = survey.Ask(q, &answers)
		if err != nil {
			fmt.Println("Add new account failed:", err)
		} else {
			u.accounts = append(u.accounts, &answers)
			u.SaveAccounts()
			goto start
		}
	} else if id < len(u.accounts) && id >= 0 {
		u.selectedAccount = u.accounts[id]
	}
}

func (u *CLI) EditConfig() {
	q := survey.Confirm{
		Message: getLanguageDes(Key_MESSAGE_EditConfig),
		Default: false,
	}

	var edit bool
	survey.AskOne(&q, &edit)
	if edit {
		b, err := json.MarshalIndent(u.config, "", "  ")
		if err != nil {
			fmt.Println("Failed to marshal current config:", err)
			return
		}
		config := string(b)
		counter := 0
	reedit:
		qe := survey.Editor{
			Message:       getLanguageDes(Key_MESSAGE_EditInEditor),
			Default:       config,
			AppendDefault: true,
			HideDefault:   true,
			FileName:      "*.json",
		}
		survey.AskOne(&qe, &config)
		var newConfig gpt.GlobalConfig
		err = json.Unmarshal([]byte(config), &newConfig)
		if err != nil && counter < 3 {
			counter++
			fmt.Println(getLanguageDes(Key_ERROR_WrongConfig))
			goto reedit
		} else if err == nil {
			u.config = &newConfig
			u.SaveConfig()
		} else if counter >= 3 {
			fmt.Println(getLanguageDes(Key_ERROR_AbortEditConfig))
		}
	}
}

func (u *CLI) SelectConversion() {
start:
	options := []any{}
	sort.Slice(u.conversions, func(i, j int) bool {
		return u.conversions[i].Created < u.conversions[j].Created
	})
	for _, conversion := range u.conversions {
		options = append(options, conversion)
	}
	addNewOption := getLanguageDes(Key_OPTION_AddNewConversion)
	options = append(options, addNewOption)
	pro := promptui.Select{
		Label: getLanguageDes(Key_LABEL_SelectConversion),
		Items: options,
	}
	id, result, err := pro.Run()
	if err != nil {
		fmt.Println("Failed to select conversion:", err)
		goto start
	}
	if result == addNewOption {
		q := []*survey.Question{
			{
				Name: "Name",
				Prompt: &survey.Input{
					Message: getLanguageDes(Key_MESSAGE_ConversionName),
				},
				Validate: survey.Required,
			},
			{
				Name: "Author",
				Prompt: &survey.Input{
					Message: getLanguageDes(Key_MESSAGE_ConversionAuthor),
				},
			},
		}
		var answers gpt.Conversion
		err = survey.Ask(q, &answers)
		if err != nil {
			fmt.Println("Add new conversion failed:", err)
			goto start
		} else {
			answers.Created = time.Now().Unix()
			answers.Updated = time.Now().Unix()
			answers.Config = u.config.DefaultConversionConfig
			u.conversions = append(u.conversions, &answers)
			u.SaveConversion(&answers)
			goto start
		}
	} else if id < len(u.conversions) && id >= 0 {
		u.selectedConversion = u.conversions[id]
	}
}

func (u *CLI) initApiConversion() {
	if u.selectedAccount == nil || u.selectedConversion == nil {
		log.Fatalln("No account or conversion selected")
	}

	cc := gpt.NewChatClient(u.selectedAccount, u.config.Host, u.config.Proxy)
	model := u.selectedConversion.Config.Model
	if strings.HasPrefix(model, "gpt-3.5-turbo") {
		model = "gpt-3.5-turbo"
	}
	tk, err := gpt.NewTiktokenCounter(model)
	if err != nil {
		log.Fatalln("Failed to init tiktoken of model["+model+"]: ", err)
	}

	ac := gpt.NewApiConversion(u.selectedConversion, tk, cc)
	u.ac = ac
}

func (u *CLI) StartMessaging() {
	if u.ac == nil {
		u.initApiConversion()
	}
	for {
		showMessages := u.ac.GetMessages()
		addMessage := getLanguageDes(Key_OPTION_AddMessage)
		sendMessage := getLanguageDes(Key_OPTION_SendMessage)
		exit := getLanguageDes(Key_OPTION_ExitConversion)
		id, result, err := u.pickOneMessage(showMessages, addMessage, sendMessage, exit)
		if err != nil {
			fmt.Printf("Failed to select message: %v\n", err)
		} else if result == addMessage {
			if role, content, err := editMessage("", ""); err != nil {
				fmt.Printf("Add message failed %v\n", err)
			} else {
				u.ac.AppendMessage(gpt.ChatRole(role), content)
			}
		} else if result == sendMessage {
			if msgs, err := u.ac.GetResponseMessages(); err != nil {
				fmt.Printf("Send failed %v\n", err)
			} else {
				id, _, err := u.pickOneMessage(msgs)
				if err != nil {
					fmt.Printf("Failed to select message: %v", err)
				}
				if id >= 0 && id < len(msgs) {
					u.ac.AppendMessage(msgs[id].Role, msgs[id].Content)
				}
			}
		} else if id >= 0 && id < len(showMessages) {
			edit := getLanguageDes(Key_OPTION_EditMessage)
			delete := getLanguageDes(Key_OPTION_DeleteMessage)
			back := getLanguageDes(Key_OPTION_Back)
			p := promptui.Select{
				Label: getLanguageDes(Key_LABEL_MessageOptions),
				Items: []string{edit, delete, back},
			}
			_, result, err := p.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			} else if result == edit {
				newRole, newContent, err := editMessage(string(showMessages[id].Role), showMessages[id].Content)
				if err != nil {
					fmt.Printf("Edit message failed %v\n", err)
				} else {
					u.ac.EditMessageAt(id, gpt.ChatRole(newRole), newContent)
				}
			} else if result == delete {
				u.ac.RemoveMessageAt(id)
			}
		} else if result == exit {
			u.SaveConversion(u.selectedConversion)
			return
		}
	}
}

func (u *CLI) pickOneMessage(msgs []gpt.ConversionMessage, extraOptions ...string) (int, string, error) {
	options := []string{}
	for i := range msgs {
		options = append(options, strconv.Itoa(i))
	}
	options = append(options, extraOptions...)
	se := survey.Select{
		Message: getLanguageDes(Key_LABEL_SelectMessage),
		Options: options,
		Default: options[len(options)-1],
		Description: func(value string, index int) string {
			if index < len(msgs) {
				raw := strings.ReplaceAll(msgs[index].String(), "\n", "")
				// charwidth := 60
				// str := insertInto(raw, charwidth, '\n')
				return raw
			}
			return ""
		},
	}

	var result string
	if err := survey.AskOne(&se, &result, survey.WithIcons(func(is *survey.IconSet) {
		is.HelpInput.Format = "Magenta"
	})); err != nil {
		return 0, "", err
	}
	id, err := strconv.Atoi(result)
	if err != nil {
		id = -1
	}
	return id, result, nil
}

func insertInto(s string, interval int, sep rune) string {
	var buffer bytes.Buffer
	before := interval - 1
	last := len(s) - 1
	for i, char := range s {
		buffer.WriteRune(char)
		if i%interval == before && i != last {
			buffer.WriteRune(sep)
		}
	}
	return buffer.String()
}

type questionMessage struct {
	Role    string
	Content string
}

func editMessage(role, content string) (string, string, error) {
	roleP := &survey.Select{
		Message: getLanguageDes(Key_MESSAGE_MessageRole),
		Options: []string{"user", "assistant"},
	}

	q := []*survey.Question{
		{
			Name:   "Role",
			Prompt: roleP,
		},
		{
			Name: "Content",
			Prompt: &survey.Editor{
				Message:       getLanguageDes(Key_MESSAGE_MessageContent),
				Default:       strings.ReplaceAll(content, "\n", " "),
				AppendDefault: true,
				HideDefault:   true,
			},
		},
	}
	if role != "" {
		roleP.Default = role
	}

	var answer questionMessage
	err := survey.Ask(q, &answer)
	if err != nil {
		return "", "", err
	}
	return answer.Role, answer.Content, nil
}

func (u *CLI) init() {
	createDirectory()
	u.accounts = loadOrCreateAccount()
	u.conversions = loadConversions()
	u.config = loadOrCreateGlobalConfig()

	u.SelectAccount()
	u.EditConfig()
	u.SelectConversion()
	u.initApiConversion()
}

func (u *CLI) Start() {
	u.StartMessaging()
}

func (u *CLI) Close() {
	u.SaveAccounts()
	u.SaveConfig()
	u.SaveConversion(u.selectedConversion)
}

func (u *CLI) SaveAccounts() {
	if err := saveAccounts(u.accounts); err != nil {
		log.Println("Failed to save accounts: ", err)
	}
}

func (u *CLI) SaveConfig() {
	if err := saveGlobalConfig(u.config); err != nil {
		log.Println("Failed to save config: ", err)
	}
}

func (u *CLI) SaveConversion(c *gpt.Conversion) {
	if err := saveConversion(c); err != nil {
		log.Println("Failed to save conversion: ", err)
	}
}
