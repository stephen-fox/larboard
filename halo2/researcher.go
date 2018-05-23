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

type Researcher struct {
	filePath string
}

func (o Researcher) IsHalo2() error {
	return o.IsMap()
}

func (o Researcher) IsMap() error {
	m, err := os.Open(o.filePath)
	if err != nil {
		return err
	}
	defer m.Close()

	raw := make([]byte, 4)

	_, err = m.ReadAt(raw, headerLocation)
	if err != nil {
		return err
	}

	toof := []int32{'t', 'o', 'o', 'f'}

	for i, c := range raw {
		if int32(c) != toof[i] {
			return errors.New(ErrorValidatingMapHeader)
		}
	}

	return nil
}

func NewResearcher(filePath string) (larboard.Researcher, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return Researcher{}, err
	}

	if info.IsDir() {
		return Researcher{}, errors.New("The specified file is a directory")
	}

	return Researcher{
		filePath: filePath,
	}, nil
}