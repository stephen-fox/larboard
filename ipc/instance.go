package ipc

import (
	"errors"
	"strings"

	"github.com/stephen-fox/larboard"
	"github.com/stephen-fox/larboard/halo2"
	"github.com/stephen-fox/larboard/internal/jsonw"
)

const (
	Halo2 Game = "halo2"
)

type Game string

type InstanceDetails struct {
	Game       Game             `json:"game"`
	Id         string           `json:"id"`
	InitialMap larboard.HaloMap `json:"initial_map"`
}

type Instance struct {
	IoOptions    IoOptions
	Details      InstanceDetails
	Cartographer larboard.Cartographer
}

func (o Instance) newErrResult(err string) Result {
	return Result{
		Id:      o.Details.Id,
		Options: o.IoOptions,
		Error:   err,
	}
}

func (o Instance) newSuccessResult(data string, message string) Result {
	return Result{
		Id:      o.Details.Id,
		Options: o.IoOptions,
		Data:    data,
		Message: message,
	}
}

func newInstance(options IoOptions, rawDetails string) (*Instance, error) {
	if len(strings.TrimSpace(rawDetails)) == 0 {
		return &Instance{}, errors.New("Please specify the InstanceDetails")
	}

	var details InstanceDetails

	err := jsonw.StringToStruct(rawDetails, &details)
	if err != nil {
		return &Instance{}, err
	}

	return NewInstance(options, details)
}

func NewInstance(options IoOptions, details InstanceDetails) (*Instance, error) {
	switch options.Source {
	case Cli:
		break
	default:
		return &Instance{}, errors.New("Unknown source: '" + string(options.Source) + "'")
	}

	if len(strings.TrimSpace(details.Id)) == 0 {
		return &Instance{}, errors.New("Please specify an instance ID")
	}

	var c larboard.Cartographer

	switch details.Game {
	case Halo2:
		cart, err := halo2.NewCartographer()
		if err != nil {
			return &Instance{}, err
		}

		c = cart
	default:
		return &Instance{}, errors.New("Unknown game: '" + string(details.Game) + "'")
	}

	instance := &Instance{
		IoOptions:    options,
		Details:      details,
		Cartographer: c,
	}

	if len(details.InitialMap.FilePath) > 0 {
		err := instance.Cartographer.SetMap(details.InitialMap)
		if err != nil {
			return &Instance{}, err
		}
	}

	return instance, nil
}
