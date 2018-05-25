package research

import (
	"github.com/stephen-fox/larboard"
	"github.com/stephen-fox/larboard/api"
)

type Data struct {
	Id         string
	Researcher larboard.Researcher
}

func IsMapValid(data Data) api.Result {
	err := data.Researcher.IsMap()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "The specified file is not a Halo map",
			Id:      data.Id,
		}
	}

	return api.Result{
		Id:      data.Id,
		Message: "This is a valid map",
	}
}

func MapName(data Data) api.Result {
	name, err := data.Researcher.Name()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "Failed to get map name",
			Id:      data.Id,
		}
	}

	return api.Result{
		Id:   data.Id,
		Data: name,
	}
}

func Scenario(data Data) api.Result {
	scenario, err := data.Researcher.Scenario()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "Failed to get map scenario",
			Id:      data.Id,
		}
	}

	return api.Result{
		Id:   data.Id,
		Data: scenario,
	}
}

func Signature(data Data) api.Result {
	signature, err := data.Researcher.Signature()
	if err != nil {
		return api.Result{
			Error:   err.Error(),
			Message: "Failed to get map's signature",
			Id:      data.Id,
		}
	}

	return api.Result{
		Id:   data.Id,
		Data: signature,
	}
}