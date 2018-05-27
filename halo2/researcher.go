package halo2

import (
	"encoding/hex"
	"errors"

	"github.com/stephen-fox/larboard"
	"github.com/stephen-fox/larboard/mapio"
)

const (
	ErrorValidatingMapHeader = "Failed to validate map header"
)

type Researcher struct {
	hmap mapio.Map
}

func (o Researcher) IsHalo2() error {
	return o.IsMap()
}

func (o Researcher) IsMap() error {
	readOptions := mapio.ReadOptions{
		Offset:  headerOffset,
		MaxSize: 4,
	}

	chunk, err := o.hmap.Read(readOptions)
	if err != nil {
		return err
	}

	toof := []int32{'t', 'o', 'o', 'f'}

	for i, c := range chunk.Bytes() {
		if int32(c) != toof[i] {
			return errors.New(ErrorValidatingMapHeader)
		}
	}

	return nil
}

func (o Researcher) Name() (string, error) {
	readOptions := mapio.ReadOptions{
		Offset:   mapNameOffset,
		MaxSize:  35,
		DoneChar: ' ',
	}

	chunk, err := o.hmap.Read(readOptions)
	if err != nil {
		return "", err
	}

	return chunk.String(), nil
}

func (o Researcher) Scenario() (string, error) {
	readOptions := mapio.ReadOptions{
		Offset:   scenarioOffset,
		MaxSize:  64,
		DoneChar: ' ',
	}

	chunk, err := o.hmap.Read(readOptions)
	if err != nil {
		return "", err
	}

	return chunk.String(), nil
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

	readOptions := mapio.ReadOptions{
		Offset:  signatureOffset,
		MaxSize: sigSize,
	}

	chunk, err := o.hmap.Read(readOptions)
	if err != nil {
		return []byte{}, err
	}

	return chunk.Bytes(), nil
}

func NewResearcher(hmap mapio.Map) (larboard.Researcher, error) {
	return Researcher{
		hmap: hmap,
	}, nil
}