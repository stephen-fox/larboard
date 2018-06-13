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
	HaloMap larboard.HaloMap
}

func (o *Cartographer) SetMap(haloMap larboard.HaloMap) error {
	o.HaloMap = haloMap

	return nil
}

func (o Cartographer) IsHalo2() error {
	return o.IsMap()
}

func (o Cartographer) IsMap() error {
	readOptions := mapio.ReadOptions{
		Offset: headerOffset,
		Size:   4,
	}

	chunk, err := mapio.AtomicRead(o.HaloMap.FilePath, readOptions)
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
		Offset: mapNameOffset,
		Size:   35,
		Until:  true,
	}

	chunk, err := mapio.AtomicRead(o.HaloMap.FilePath, readOptions)
	if err != nil {
		return "", err
	}

	return chunk.String(), nil
}

func (o Cartographer) Scenario() (string, error) {
	readOptions := mapio.ReadOptions{
		Offset: scenarioOffset,
		Size:   64,
		Until:  true,
	}

	chunk, err := mapio.AtomicRead(o.HaloMap.FilePath, readOptions)
	if err != nil {
		return "", err
	}

	return chunk.String(), nil
}

func (o Cartographer) Signature() (string, error) {
	readOptions := mapio.ReadOptions{
		Offset: signatureOffset,
		Size:   4,
	}

	chunk, err := mapio.AtomicRead(o.HaloMap.FilePath, readOptions)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(chunk.Bytes()), nil
}

func (o Cartographer) Sign() (string, error) {
	openOptions := mapio.OpenOptions{
		AllowWrite: true,
	}

	f, err := mapio.Open(o.HaloMap.FilePath, openOptions)
	if err != nil {
		return "", err
	}
	defer f.Close()

	currentOffset := int64(bodyOffset)
	var bodyXor [4]byte
	const buffSize = 16384

	for {
		ro := mapio.ReadOptions{
			Offset: currentOffset,
			Size:   buffSize,
		}

		currentOffset = currentOffset + buffSize

		chunk, err := mapio.Read(f, ro)
		if err != nil {
			return "", err
		}

		for i, b := range chunk.Data {
			bodyXor[i&3] ^= b
		}

		if chunk.Eof {
			break
		}
	}

	wo := mapio.WriteOptions{
		Offset: signatureOffset,
		Data:   bodyXor[:],
	}

	err = mapio.Write(f, wo)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(wo.Data), nil
}

func NewCartographer() (larboard.Cartographer, error) {
	return &Cartographer{}, nil
}