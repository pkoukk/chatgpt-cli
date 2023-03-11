package gpt

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ChatClient struct {
	account Account
	path    string
	client  *resty.Client
}

func NewChatClient(account Account, host string, proxy string) *ChatClient {
	client := resty.New()
	if proxy != "" {
		client.SetProxy(proxy)
	}
	return &ChatClient{
		account: account,
		path:    host + "/v1/chat/completions",
		client:  client,
	}
}

type ChatRequest struct {
	Model       string         `json:"model"`
	Messages    []*ChatMessage `json:"messages"`
	Temperature float64        `json:"temperature,omitempty"`
	TopP        float64        `json:"top_p,omitempty"`
	N           int            `json:"n,omitempty"`
	Stream      bool           `json:"stream,omitempty"`
	Stop        string         `json:"stop,omitempty"`
	MaxTokens   string         `json:"max_tokens,omitempty"`
	User        string         `json:"user,omitempty"`
}

type ChatRole string

const (
	CHAT_ROLE_SYSTEM    ChatRole = "system"
	CHAT_ROLE_USER      ChatRole = "user"
	CHAT_ROLE_ASSISTANT ChatRole = "assistant"
)

type ChatMessage struct {
	Content string   `json:"content"`
	Role    ChatRole `json:"role"`
}

func (c *ChatMessage) String() string {
	return string(c.Role) + ": " + c.Content
}

type ChatResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Usage   *Usage    `json:"usage"`
	Choices []*Choice `json:"choices"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
}

func (cc *ChatClient) Request(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	fmt.Println(req.Messages)
	var res ChatResponse
	_, err := cc.client.R().
		SetAuthToken(cc.account.Token()).
		SetContext(ctx).
		SetBody(req).
		SetResult(&res).
		Post(cc.path)
	if err != nil {
		return nil, err
	}
	if len(res.Choices) == 0 {
		return nil, errors.New("no choices")
	}
	return &res, nil
}
