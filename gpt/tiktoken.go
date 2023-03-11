package gpt

import "github.com/pkoukk/tiktoken-go"

type TokenCounter interface {
	CountTokens(string) int
}

type TiktokenCounter struct {
	tk *tiktoken.Tiktoken
}

func NewTiktokenCounter(model string) (*TiktokenCounter, error) {
	tk, errr := tiktoken.EncodingForModel(model)
	if errr != nil {
		return nil, errr
	}
	return &TiktokenCounter{
		tk: tk,
	}, nil
}

func (t *TiktokenCounter) CountTokens(text string) int {
	return len(t.tk.Encode(text, nil, nil))
}
