package halo2

import (
	"errors"
	"os"

	"github.com/stephen-fox/larboard"
)

const (
	ErrorValidatingMapHeader = "Failed to validate map header"

	mapNameOffset   = 408
	scenarioOffset  = 444
	signatureOffset = 720
	headerOffset    = 2044
)

type Researcher struct {
	m *os.File
}

func (o Researcher) IsHalo2() error {
	return o.IsMap()
}

func (o Researcher) IsMap() error {
	fc := larboard.FileChunk{
		Offset:  headerOffset,
		MaxSize: 4,
	}

	raw, err := fc.Bytes(o.m)
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
	fc := larboard.FileChunk{
		Offset:   mapNameOffset,
		MaxSize:  35,
		DoneChar: ' ',
	}

	return fc.String(o.m)
}

func (o Researcher) Scenario() (string, error) {
	fc := larboard.FileChunk{
		Offset:   scenarioOffset,
		MaxSize:  64,
		DoneChar: ' ',
	}

	return fc.String(o.m)
}

func (o Researcher) Signature() (string, error) {
	fc := larboard.FileChunk{
		Offset:   signatureOffset,
		MaxSize:  4,
	}

	return fc.String(o.m)
}

func NewResearcher(m *os.File) (larboard.Researcher, error) {
	return Researcher{
		m: m,
	}, nil
}