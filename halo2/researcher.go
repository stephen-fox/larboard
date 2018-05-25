package halo2

import (
	"encoding/hex"
	"errors"
	"os"

	"github.com/stephen-fox/larboard"
	"github.com/stephen-fox/larboard/chunking"
)

const (
	ErrorValidatingMapHeader = "Failed to validate map header"
)

type Researcher struct {
	m *os.File
}

func (o Researcher) IsHalo2() error {
	return o.IsMap()
}

func (o Researcher) IsMap() error {
	fc := chunking.FileChunk{
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
	fc := chunking.FileChunk{
		Offset:   mapNameOffset,
		MaxSize:  35,
		DoneChar: ' ',
	}

	return fc.String(o.m)
}

func (o Researcher) Scenario() (string, error) {
	fc := chunking.FileChunk{
		Offset:   scenarioOffset,
		MaxSize:  64,
		DoneChar: ' ',
	}

	return fc.String(o.m)
}

func (o Researcher) Signature() (string, error) {
	raw, err := o.SignatureRaw()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(raw), nil
}

func (o Researcher) SignatureRaw() ([]byte, error) {
	sigSize := 4

	fc := chunking.FileChunk{
		Offset:  signatureOffset,
		MaxSize: sigSize,
	}

	raw, err := fc.Bytes(o.m)
	if err != nil {
		return raw, err
	}

	return raw, nil
}

func NewResearcher(m *os.File) (larboard.Researcher, error) {
	return Researcher{
		m: m,
	}, nil
}