package halo2

import (
	"errors"
	"os"

	"github.com/stephen-fox/larboard"
)

const (
	ErrorValidatingMapHeader = "Failed to validate map header"

	headerLocation = 2044
)

type MapDescriber struct {
	filePath string
}

func (o MapDescriber) IsHalo2() (bool, error) {
	return o.IsMap()
}

func (o MapDescriber) IsMap() (bool, error) {
	m, err := os.Open(o.filePath)
	if err != nil {
		return false, err
	}
	defer m.Close()

	raw := make([]byte, 4)

	_, err = m.ReadAt(raw, headerLocation)
	if err != nil {
		return false, err
	}

	toof := []int32{'t', 'o', 'o', 'f'}

	for i, c := range raw {
		if int32(c) != toof[i] {
			return false, errors.New(ErrorValidatingMapHeader)
		}
	}

	return true, nil
}

func NewMap(filePath string) (larboard.MapDescriber, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return MapDescriber{}, err
	}

	if info.IsDir() {
		return MapDescriber{}, errors.New("The specified file is a directory")
	}

	return MapDescriber{
		filePath: filePath,
	}, nil
}