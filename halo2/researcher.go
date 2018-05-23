package halo2

import (
	"errors"
	"os"

	"github.com/stephen-fox/larboard"
)

const (
	ErrorValidatingMapHeader = "Failed to validate map header"

	mapNameOffset  = 408
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

func (o Researcher) Name() (string, error) {
	fcs := larboard.FileChunkStringer{
		FilePath: o.filePath,
		Offset:   mapNameOffset,
		MaxSize:  35,
		DoneChar: ' ',
	}

	return fcs.Read()
}

func (o Researcher) Scenario() (string, error) {
	fcs := larboard.FileChunkStringer{
		FilePath: o.filePath,
		Offset:   scenarioOffset,
		MaxSize:  64,
		DoneChar: ' ',
	}

	return fcs.Read()
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