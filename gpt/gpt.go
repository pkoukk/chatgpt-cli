package gpt

import (
	"fmt"
	"time"
)

type Account interface {
	Name() string
	Token() string
}

type Conversion struct {
	Name     string
	Author   string
	Created  int64
	Updated  int64
	Config   ConversionConfig
	Messages []*ConversionMessage
}

type ConversionMessage struct {
	*ChatMessage
	TokenCount int
}

func (c *Conversion) String() string {
	return fmt.Sprintf("<%s> - %s | %s", c.Name, c.Author, time.Unix(c.Created, 0).Format("2006-01-02"))
}
