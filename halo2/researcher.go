package halo2

import (
	"errors"
	"os"

	"github.com/stephen-fox/larboard"
)

const (
	ErrorValidatingMapHeader = "Failed to validate map header"

	scenarioOffset = 444
	headerOffset   = 2044
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

	_, err = m.ReadAt(raw, headerOffset)
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

func (o Researcher) Scenario() (string, error) {
	m, err := os.Open(o.filePath)
	if err != nil {
		return "", err
	}
	defer m.Close()

	raw := make([]byte, 64)

	_, err = m.ReadAt(raw, scenarioOffset)
	if err != nil {
		return "", err
	}

	var chars []int32

	for _, r := range raw {
		c := int32(r)

		if c == ' ' {
			break
		}

		chars = append(chars, c)
	}

	return string(chars), nil
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