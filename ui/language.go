package ui

import (
	"os"
	"strings"
)

func init() {
	lang := os.Getenv("LANG")
	switch {
	case strings.EqualFold(lang, "en"):
		initEnUs()
	default:
		initZhCn()
	}
}

func getLanguageDes(key LanguagePackKey) string {
	return LanguagePack[key]
}

type LanguagePackKey string

const (
	Key_LABEL_SelectAccount         LanguagePackKey = "select_account"
	Key_LABEL_SelectConversion      LanguagePackKey = "select_conversion"
	Key_LABEL_SelectMessage         LanguagePackKey = "select_message"
	Key_LABEL_SelectResponseMessage LanguagePackKey = "select_response_message"
	Key_LABEL_MessageOptions        LanguagePackKey = "message_options"

	Key_OPTION_AddNewAccount    LanguagePackKey = "add_new_account"
	Key_OPTION_AddNewConversion LanguagePackKey = "add_new_conversion"
	Key_OPTION_AddMessage       LanguagePackKey = "add_message"
	Key_OPTION_SendMessage      LanguagePackKey = "send_message"
	Key_OPTION_ExitConversion   LanguagePackKey = "exit_conversion"
	Key_OPTION_EditMessage      LanguagePackKey = "edit_message"
	Key_OPTION_DeleteMessage    LanguagePackKey = "delete_message"
	Key_OPTION_Back             LanguagePackKey = "back"

	Key_MESSAGE_ConversionName   LanguagePackKey = "conversion_name"
	Key_MESSAGE_ConversionAuthor LanguagePackKey = "conversion_author"
	Key_MESSAGE_MessageRole      LanguagePackKey = "message_role"
	Key_MESSAGE_MessageContent   LanguagePackKey = "message_content"
	Key_MESSAGE_AccountName      LanguagePackKey = "account_name"
	Key_MESSAGE_SecretKey        LanguagePackKey = "secret_key"
	Key_MESSAGE_EditConfig       LanguagePackKey = "edit_config"
	Key_MESSAGE_EditInEditor     LanguagePackKey = "edit_in_editor"

	Key_ERROR_WrongConfig     LanguagePackKey = "wrong_config"
	Key_ERROR_AbortEditConfig LanguagePackKey = "abort_edit_config"
)

var LanguagePack map[LanguagePackKey]string

func initEnUs() {
	LanguagePack = map[LanguagePackKey]string{
		Key_LABEL_SelectAccount:         "Select account >> ",
		Key_LABEL_SelectConversion:      "Select conversion >> ",
		Key_LABEL_SelectMessage:         "Select message >> ",
		Key_LABEL_SelectResponseMessage: "Pick one response message you like to continue >> ",
		Key_LABEL_MessageOptions:        "What you want  >> ",

		Key_OPTION_AddNewAccount:    "Add new account",
		Key_OPTION_AddNewConversion: "Add new conversion",
		Key_OPTION_AddMessage:       "Add message",
		Key_OPTION_SendMessage:      "Send message",
		Key_OPTION_ExitConversion:   "Exit conversion",
		Key_OPTION_EditMessage:      "Edit message",
		Key_OPTION_DeleteMessage:    "Delete message",
		Key_OPTION_Back:             "Back",

		Key_MESSAGE_ConversionName:   "Conversion name:",
		Key_MESSAGE_ConversionAuthor: "Conversion author:",
		Key_MESSAGE_MessageRole:      "Message role:",
		Key_MESSAGE_MessageContent:   "Message content:",
		Key_MESSAGE_AccountName:      "Account name:",
		Key_MESSAGE_SecretKey:        "Secret key(https://platform.openai.com/account/api-keys):",
		Key_MESSAGE_EditConfig:       "Do you want edit config?",
		Key_MESSAGE_EditInEditor:     "Press Enter to edit config in following editor,save and exit to continue",

		Key_ERROR_WrongConfig:     "Not a valid config,please check",
		Key_ERROR_AbortEditConfig: "Invalid edit for 3 times,abort edit config",
	}
}

func initZhCn() {
	LanguagePack = map[LanguagePackKey]string{
		Key_LABEL_SelectAccount:         "???????????? >> ",
		Key_LABEL_SelectConversion:      "???????????? >> ",
		Key_LABEL_SelectMessage:         "???????????? >> ",
		Key_LABEL_SelectResponseMessage: "???????????????????????????????????????????????? >>",
		Key_LABEL_MessageOptions:        "???????????? >> ",

		Key_OPTION_AddNewAccount:    "???????????????",
		Key_OPTION_AddNewConversion: "???????????????",
		Key_OPTION_AddMessage:       "????????????",
		Key_OPTION_SendMessage:      "????????????",
		Key_OPTION_ExitConversion:   "????????????",
		Key_OPTION_EditMessage:      "????????????",
		Key_OPTION_DeleteMessage:    "????????????",
		Key_OPTION_Back:             "??????",

		Key_MESSAGE_ConversionName:   "????????????:",
		Key_MESSAGE_ConversionAuthor: "????????????:",
		Key_MESSAGE_MessageRole:      "????????????:",
		Key_MESSAGE_MessageContent:   "????????????:",
		Key_MESSAGE_AccountName:      "????????????:",
		Key_MESSAGE_SecretKey:        "??????(https://platform.openai.com/account/api-keys):",
		Key_MESSAGE_EditConfig:       "???????????????????",
		Key_MESSAGE_EditInEditor:     "??????<?????????>???????????????????????????,????????????????????????",

		Key_ERROR_WrongConfig:     "????????????,?????????",
		Key_ERROR_AbortEditConfig: "??????????????????3???,????????????",
	}
}
