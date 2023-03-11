package gpt

import "context"

type ApiConversion struct {
	c             *Conversion
	tokenCounter  TokenCounter
	chatClient    *ChatClient
	messageWindow int
}

func NewApiConversion(c *Conversion, tokenCounter TokenCounter, client *ChatClient) *ApiConversion {
	for i := range c.Messages {
		if c.Messages[i].TokenCount == 0 {
			c.Messages[i].TokenCount = tokenCounter.CountTokens(c.Messages[i].Content)
		}
	}

	ac := &ApiConversion{
		c:             c,
		tokenCounter:  tokenCounter,
		messageWindow: 0,
		chatClient:    client,
	}
	ac.calcWindow()
	return ac
}

func (ac *ApiConversion) calcWindow() {
	maxToken := ac.c.Config.MaxTokens
	if maxToken == 0 {
		maxToken = 2048
	}
	window := 0
	for i := len(ac.c.Messages) - 1; i >= 0; i-- {
		window = i
		// +4 for role cost
		maxToken -= (ac.c.Messages[i].TokenCount + 4)
		if maxToken <= 0 {
			break
		}
	}
	ac.messageWindow = window
}

func (ac *ApiConversion) GetConversion() Conversion {
	return *ac.c
}

func (ac *ApiConversion) GetMessages() []ConversionMessage {
	var messages []ConversionMessage
	for i := ac.messageWindow; i < len(ac.c.Messages); i++ {
		messages = append(messages, *ac.c.Messages[i])
	}
	return messages
}

func (ac *ApiConversion) EditMessageAt(index int, role ChatRole, content string) {
	if index < 0 || index >= len(ac.c.Messages) {
		return
	}
	ac.c.Messages[index].Content = content
	ac.c.Messages[index].Role = role
	ac.c.Messages[index].TokenCount = ac.tokenCounter.CountTokens(content)
	ac.calcWindow()

}

func (ac *ApiConversion) AppendMessage(role ChatRole, content string) {
	ac.c.Messages = append(ac.c.Messages, &ConversionMessage{
		ChatMessage: &ChatMessage{Role: role, Content: content},
		TokenCount:  ac.tokenCounter.CountTokens(content),
	})
	ac.calcWindow()
}

func (ac *ApiConversion) AddMessageAt(role ChatRole, content string, index int) {
	if index < 0 || index > len(ac.c.Messages) {
		return
	}
	ac.c.Messages = append(ac.c.Messages[:index], append([]*ConversionMessage{{
		ChatMessage: &ChatMessage{Role: role, Content: content},
		TokenCount:  ac.tokenCounter.CountTokens(content),
	}}, ac.c.Messages[index:]...)...)
	ac.calcWindow()
}

func (ac *ApiConversion) RemoveMessageAt(index int) {
	if index < 0 || index >= len(ac.c.Messages) {
		return
	}
	ac.c.Messages = append(ac.c.Messages[:index], ac.c.Messages[index+1:]...)
	ac.calcWindow()
}

func (ac *ApiConversion) GetResponseMessages() ([]*ConversionMessage, error) {
	windowedMessages := ac.c.Messages[ac.messageWindow:]
	chatMessages := make([]*ChatMessage, len(windowedMessages))
	for i := range windowedMessages {
		chatMessages[i] = windowedMessages[i].ChatMessage
	}

	req := &ChatRequest{
		Model:       ac.c.Config.Model,
		Messages:    chatMessages,
		Temperature: ac.c.Config.Temperature,
		TopP:        ac.c.Config.TopP,
		N:           ac.c.Config.N,
	}
	res, err := ac.chatClient.Request(context.Background(), req)
	if err != nil {
		return nil, err
	}

	messages := make([]*ConversionMessage, len(res.Choices))
	for i := range res.Choices {
		if res.Choices[i].Message.Content != "" {
			messages[i] = &ConversionMessage{
				ChatMessage: &res.Choices[i].Message,
			}
		}
	}

	return messages, nil
}
