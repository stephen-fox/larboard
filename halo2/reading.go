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

type Cartographer struct {
	mapper mapio.Mapper
}

func (o *Cartographer) SetMap(haloMap larboard.HaloMap) error {
	o.mapper.FilePath = haloMap.FilePath

	return nil
}

func (o Cartographer) IsHalo2() error {
	return o.IsMap()
}

func (o Cartographer) IsMap() error {
	readOptions := mapio.ReadOptions{
		Offset:  headerOffset,
		MaxSize: 4,
	}

	chunk, err := o.mapper.Read(readOptions)
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

func (o Cartographer) Name() (string, error) {
	readOptions := mapio.ReadOptions{
		Offset:  mapNameOffset,
		MaxSize: 35,
		Stop:    true,
	}

	chunk, err := o.mapper.Read(readOptions)
	if err != nil {
		return "", err
	}

	return chunk.String(), nil
}

func (o Cartographer) Scenario() (string, error) {
	readOptions := mapio.ReadOptions{
		Offset:  scenarioOffset,
		MaxSize: 64,
		Stop:    true,
	}

	chunk, err := o.mapper.Read(readOptions)
	if err != nil {
		return "", err
	}

	return chunk.String(), nil
}

func (o Cartographer) Signature() (string, error) {
	readOptions := mapio.ReadOptions{
		Offset:  signatureOffset,
		MaxSize: 4,
	}

	chunk, err := o.mapper.Read(readOptions)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(chunk.Bytes()), nil
}

func (o Cartographer) Sign() (string, error) {
	return "", errors.New("Not implmented yet")
}

func NewCartographer() (larboard.Cartographer, error) {
	return &Cartographer{}, nil
}