package ui

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/pkoukk/chatgpt-cli/gpt"
)

func createDirectory() {
	if err := os.MkdirAll("./chatgpt-cli", os.ModePerm); err != nil {
		log.Fatalln("cannot create work directory:", err)
	}
	if err := os.MkdirAll("./chatgpt-cli/conversions", os.ModePerm); err != nil {
		log.Fatalln("cannot create conversions directory:", err)
	}
}

func loadOrCreateAccount() []*gpt.SecretKeyAccount {
	if _, err := os.Stat("./chatgpt-cli/accounts.json"); err == nil {
		contents, err := ioutil.ReadFile("./chatgpt-cli/accounts.json")
		if err != nil {
			panic(err)
		}
		if len(contents) != 0 {
			var accounts []*gpt.SecretKeyAccount
			if err := json.Unmarshal(contents, &accounts); err != nil {
				log.Fatalln("Error unmarshalling accounts.json:", err)
			}
			return accounts
		}
	} else if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create("./chatgpt-cli/accounts.json")
		if err != nil {
			log.Fatalln("cannot create accounts.json:", err)
		}
	} else {
		log.Fatalln("cannot open accounts.json:", err)
	}

	return []*gpt.SecretKeyAccount{}
}

func loadOrCreateGlobalConfig() *gpt.GlobalConfig {
	if _, err := os.Stat("./chatgpt-cli/config.json"); err == nil {
		contents, err := ioutil.ReadFile("./chatgpt-cli/config.json")
		if err != nil {
			panic(err)
		}
		if len(contents) != 0 {
			var config gpt.GlobalConfig
			if err := json.Unmarshal(contents, &config); err != nil {
				log.Fatalln("Error unmarshalling config.json:", err)
			}
			return &config
		}
	} else if errors.Is(err, os.ErrNotExist) {

	} else {
		log.Fatalln("cannot open config.json:", err)
	}

	return &gpt.GlobalConfig{
		Host: "https://api.openai.com",
		DefaultConversionConfig: gpt.ConversionConfig{
			Model:     "gpt-3.5-turbo",
			N:         2,
			MaxTokens: 2048,
		},
	}
}

func loadConversions() []*gpt.Conversion {
	de, err := os.ReadDir("./chatgpt-cli/conversions")
	if err != nil {
		log.Fatalln("cannot read conversions directory:", err)
	}
	var conversions []*gpt.Conversion
	for _, file := range de {
		if file.IsDir() {
			continue
		}
		contents, err := ioutil.ReadFile("./chatgpt-cli/conversions/" + file.Name())
		if err != nil {
			log.Fatalln("cannot read conversion file:", err)
		}
		var conversion gpt.Conversion
		if err := json.Unmarshal(contents, &conversion); err != nil {
			log.Fatalln("Error unmarshalling conversion:", err)
		}
		if conversion.Config.Model == "" {
			conversion.Config.Model = "gpt-3.5-turbo"
		}
		if conversion.Config.Base64 {
			for i := range conversion.Messages {
				b, e := base64.StdEncoding.DecodeString(conversion.Messages[i].Content)
				if e != nil {
					log.Fatalln("Error decoding base64:", e)
				}
				conversion.Messages[i].Content = string(b)
			}
		}
		conversions = append(conversions, &conversion)
	}
	return conversions
}

func saveConversions(conversions []*gpt.Conversion) error {
	for _, conversion := range conversions {
		if err := saveConversion(*conversion); err != nil {
			return err
		}
	}
	return nil
}

func saveConversion(conversion gpt.Conversion) error {
	conversion.Updated = time.Now().Unix()
	fileName := "./chatgpt-cli/conversions/" + conversion.Name + ".json"
	if conversion.Config.Base64 {
		for i := range conversion.Messages {
			c1 := base64.StdEncoding.EncodeToString([]byte(conversion.Messages[i].Content))
			conversion.Messages[i].Content = c1
		}
	}
	content, err := json.MarshalIndent(conversion, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal conversion %s: %w", conversion.Name, err)
	}
	if err := os.WriteFile(fileName, content, 0644); err != nil {
		return fmt.Errorf("cannot write conversion %s: %w", conversion.Name, err)
	}
	return nil
}

func saveGlobalConfig(config *gpt.GlobalConfig) error {
	content, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal config: %w", err)
	}
	if err := os.WriteFile("./chatgpt-cli/config.json", content, 0644); err != nil {
		return fmt.Errorf("cannot write config: %w", err)
	}
	return nil
}

func saveAccounts(accounts []*gpt.SecretKeyAccount) error {
	content, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal accounts: %w", err)
	}
	if err := os.WriteFile("./chatgpt-cli/accounts.json", content, 0644); err != nil {
		return fmt.Errorf("cannot write accounts: %w", err)
	}
	return nil
}
