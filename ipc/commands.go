package ipc

import (
	"strings"

	"github.com/stephen-fox/larboard"
	"github.com/stephen-fox/larboard/internal/jsonw"
)

const (
	newCommand   = "new"
	apiSeparator = "|"
)

var (
	stringsToCommandFuncs = map[string]commandFunc{
		newCommand:  newI,
		"setmap":    setMap,
		"ishalo2":   isHalo2,
		"ismap":     isMap,
		"name":      name,
		"scenario":  scenario,
		"signature": signature,
		"sign":      sign,
	}
)

type commandFunc func(command Command, instance *Instance) Result

type Command struct {
	IsUnknown  bool
	HasArgs    bool
	Func       commandFunc
	Name       string
	Args       string
	InstanceId string
}

type Result struct {
	Data    string    `json:"data"`
	Error   string    `json:"error"`
	Id      string    `json:"id"`
	Message string    `json:"message"`
	Options IoOptions `json:"-"`
}

func (o Result) IsError() bool {
	if len(strings.TrimSpace(o.Error)) > 0 {
		return true
	}

	return false
}

func (o Result) FormatOutput() string {
	switch o.Options.Source {
	case Cli:
		fallthrough
	default:
		out := o.Message

		if o.Options.HumanReadableOutput {
			raw, err := jsonw.ToPrettyString(o)
			if err == nil {
				out = string(raw)
			}
		} else {
			conv, err := jsonw.ToString(o)
			if err == nil {
				out = conv
			}
		}

		return out
	}
}

func newCliCommand(rawInput string) (Command) {
	idAndMethod := strings.Split(rawInput, apiSeparator)

	switch len(idAndMethod) {
	case 0:
		return Command{}
	case 1:
		isValid, command := getCommand(idAndMethod[0])
		if isValid {
			return command
		}
	default:
		isValid, command := getCommand(idAndMethod[1])
		if isValid {
			command.InstanceId = idAndMethod[0]
			return command
		}
	}

	return Command{
		Name:      rawInput,
		IsUnknown: true,
		Func:      defaultMethod,
	}
}

func getCommand(methodAndArgs string) (bool, Command) {
	methodRaw := methodAndArgs
	args := ""

	if strings.Contains(methodAndArgs, "{") && strings.HasSuffix(methodAndArgs, "}") {
		sep := "{"
		parts := strings.Split(methodAndArgs, sep)
		methodRaw = parts[0]
		args = sep + strings.Join(parts[1:], sep)
	}

	methodRaw = strings.ToLower(methodRaw)

	cFunc, ok := stringsToCommandFuncs[methodRaw]
	if ok {
		return true, Command{
			Name:    methodRaw,
			Args:    args,
			HasArgs: len(args) > 0,
			Func:    cFunc,
		}
	}

	return false, Command{
		Name: methodRaw,
		Func: defaultMethod,
	}
}

func defaultMethod(command Command, instance *Instance) Result {
	return Result{}
}

func newI(command Command, instance *Instance) Result {
	return Result{}
}

func setMap(command Command, instance *Instance) Result {
	if command.HasArgs {
		var haloMap larboard.HaloMap

		err := jsonw.StringToStruct(command.Args, &haloMap)
		if err != nil {
			return instance.newErrResult(err.Error())
		}

		err = instance.Cartographer.SetMap(haloMap)
		if err != nil {
			return instance.newErrResult(err.Error())
		}

		return instance.newSuccessResult("", "")
	}

	return instance.newErrResult("Please specify a HaloMap")
}

func isHalo2(command Command, instance *Instance) Result {
	err := instance.Cartographer.IsHalo2()
	if err != nil {
		return instance.newErrResult(err.Error())
	}

	return instance.newSuccessResult("", "")
}

func isMap(command Command, instance *Instance) Result {
	err := instance.Cartographer.IsMap()
	if err != nil {
		return instance.newErrResult(err.Error())
	}

	return instance.newSuccessResult("", "")
}

func name(command Command, instance *Instance) Result {
	name, err := instance.Cartographer.Name()
	if err != nil {
		return instance.newErrResult(err.Error())
	}

	return instance.newSuccessResult(name, "")
}

func scenario(command Command, instance *Instance) Result {
	scenario, err := instance.Cartographer.Scenario()
	if err != nil {
		return instance.newErrResult(err.Error())
	}

	return instance.newSuccessResult(scenario, "")
}

func signature(command Command, instance *Instance) Result {
	signature, err := instance.Cartographer.Signature()
	if err != nil {
		return instance.newErrResult(err.Error())
	}

	return instance.newSuccessResult(signature, "")
}

func sign(command Command, instance *Instance) Result {
	newSignature, err := instance.Cartographer.Sign()
	if err != nil {
		return instance.newErrResult(err.Error())
	}

	return instance.newSuccessResult(newSignature, "")
}
