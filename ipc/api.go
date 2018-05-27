package ipc

import (
	"strings"
	"github.com/stephen-fox/larboard"
)

const (
	SingleRunCli Source = "cli"
)

const (
	General  Class = "general"
	Research Class = "research"
)

const (
	Open      Method = "open"
	Close     Method = "close"
	IsHalo2   Method = "is_halo_2"
	IsValid   Method = "is_valid"
	Name      Method = "name"
	Scenario  Method = "scenario"
	Signature Method = "signature"
)

type Source string

type Class string

type Method string

type Instance struct {
	Researcher   larboard.Researcher
	Cartographer larboard.Cartographer
}

type Input struct {
	Id     string   `json:"id"`
	Source Source   `json:"-"`
	Class  Class    `json:"class"`
	Method Method   `json:"method"`
	Args   []string `json:"args"`
}

func (o Input) Execute() Result {
	switch o.Class {
	case General:
		switch o.Method {
		case Open:

		}
	}
}

type Result struct {
	Data    string `json:"data"`
	Error   string `json:"error"`
	Id      string `json:"id"`
	Message string `json:"message"`
	Source  Source   `json:"-"`
}

func (o Result) IsError() bool {
	if len(strings.TrimSpace(o.Error)) > 0 {
		return true
	}

	return false
}

func (o Result) FormatOutput() string {
	switch o.Source {
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

func NewInstance() (Instance, error) {
	
}