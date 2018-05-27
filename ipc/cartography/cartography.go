package cartography

import (
	"github.com/stephen-fox/larboard"
	"github.com/stephen-fox/larboard/ipc/api"
)

type Data struct {
	Id           string
	Cartographer larboard.Cartographer
}

func Sign(data api.Input) api.Result {
	newSignature, err := data.Cartographer.Sign()
	if err != nil {
		return api.Result{
			Id:      data.Id,
			Error:   err.Error(),
			Message: "Failed to sign map",
		}
	}

	return api.Result{
		Id:      data.Id,
		Data:    newSignature,
		Message: "Signed map",
	}
}