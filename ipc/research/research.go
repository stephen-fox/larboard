package research

import (
	"github.com/stephen-fox/larboard/ipc/api"
)

func IsMapValid(input api.Input) api.Result {
	err := input.Researcher.IsMap()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "The specified file is not a Halo map",
			Id:      input.Id,
		}
	}

	return api.Result{
		Id:      input.Id,
		Message: "This is a valid map",
		Data:    "true",
	}
}

func MapName(input api.Input) api.Result {
	name, err := input.Researcher.Name()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "Failed to get map name",
			Id:      input.Id,
		}
	}

	return api.Result{
		Id:   input.Id,
		Data: name,
	}
}

func Scenario(input api.Input) api.Result {
	scenario, err := input.Researcher.Scenario()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "Failed to get map scenario",
			Id:      input.Id,
		}
	}

	return api.Result{
		Id:   input.Id,
		Data: scenario,
	}
}

func Signature(input api.Input) api.Result {
	signature, err := input.Researcher.Signature()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "Failed to get map's signature",
			Id:      input.Id,
		}
	}

	return api.Result{
		Id:   input.Id,
		Data: signature,
	}
}