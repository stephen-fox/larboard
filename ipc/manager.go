package ipc

import (
	"errors"
	"strings"
)

const (
	Cli Source = "cli"
)

type Source string

type IoOptions struct {
	HumanReadableOutput bool
	Source              Source
}

type Manager struct {
	IdsToInstances map[string]*Instance
}

func (o *Manager) AddInstance(instance *Instance) error {
	_, ok := o.IdsToInstances[instance.Details.Id]
	if ok {
		return errors.New("The specified instance already exists")
	}

	o.IdsToInstances[instance.Details.Id] = instance

	return nil
}

func (o *Manager) BlockAndParseInput(options IoOptions, rawInputs chan string, results chan Result) error {
	for raw := range rawInputs {
		if len(strings.TrimSpace(raw)) == 0 {
			continue
		}

		command := newCliCommand(raw)

		if command.IsUnknown {
			results <- Result{
				Error: "Unknown command: '" + command.Name + "'",
			}

			continue
		}

		if command.Name == newCommand {
			instance, err := newInstance(options, command.Args)
			if err != nil {
				results <- Result{
					Error: err.Error(),
				}

				continue
			}

			o.IdsToInstances[instance.Details.Id] = instance
		} else {
			if len(strings.TrimSpace(command.InstanceId)) == 0 && len(o.IdsToInstances) == 1 {
				for _, in := range o.IdsToInstances {
					command.InstanceId = in.Details.Id
				}
			}

			instance, ok := o.IdsToInstances[command.InstanceId]
			if ok {
				results <- command.Func(command, instance)
				continue
			}

			results <- Result{
				Error: "Unknown instance: '" + command.InstanceId + "'",
			}
		}
	}

	return nil
}

func NewManager() (Manager, error) {
	return Manager{
		IdsToInstances: make(map[string]*Instance),
	}, nil
}
