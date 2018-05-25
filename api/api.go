package api

import (
	"strings"
)

const (
	SingleRunCli Type = "cli"
)

type Type string

type Result struct {
	Data    string `json:"data"`
	Error   string `json:"error"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

func (o Result) IsError() bool {
	if len(strings.TrimSpace(o.Error)) > 0 {
		return true
	}

	return false
}

func (o Result) FormatOutput(apiType Type) string {
	switch apiType {
	case SingleRunCli:
		fallthrough
	default:
		out := o.Message

		if len(out) > 0 && len(o.Error) > 0 && len(o.Message) > 0 {
			out = out + " - "
		}

		if len(o.Error) == 0 {
			out = out + o.Data
		} else {
			out = out + o.Error
		}

		return out
	}
}