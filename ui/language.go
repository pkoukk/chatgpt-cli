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
		Key_LABEL_SelectAccount:         "选择账号 >> ",
		Key_LABEL_SelectConversion:      "选择会话 >> ",
		Key_LABEL_SelectMessage:         "选择消息 >> ",
		Key_LABEL_SelectResponseMessage: "挑选一条你喜欢的回复添加到会话中 >>",
		Key_LABEL_MessageOptions:        "选择操作 >> ",

		Key_OPTION_AddNewAccount:    "添加新账号",
		Key_OPTION_AddNewConversion: "添加新会话",
		Key_OPTION_AddMessage:       "添加消息",
		Key_OPTION_SendMessage:      "发送消息",
		Key_OPTION_ExitConversion:   "退出会话",
		Key_OPTION_EditMessage:      "编辑消息",
		Key_OPTION_DeleteMessage:    "删除消息",
		Key_OPTION_Back:             "返回",

		Key_MESSAGE_ConversionName:   "会话名称:",
		Key_MESSAGE_ConversionAuthor: "会话作者:",
		Key_MESSAGE_MessageRole:      "消息角色:",
		Key_MESSAGE_MessageContent:   "消息内容:",
		Key_MESSAGE_AccountName:      "账号名称:",
		Key_MESSAGE_SecretKey:        "密钥(https://platform.openai.com/account/api-keys):",
		Key_MESSAGE_EditConfig:       "是否编辑配置?",
		Key_MESSAGE_EditInEditor:     "按下<回车键>在编辑器中编辑配置,完成后保存并退出",

		Key_ERROR_WrongConfig:     "配置错误,请检查",
		Key_ERROR_AbortEditConfig: "编辑配置错误3次,放弃编辑",
	}
}
