package mapio

import (
	"errors"
	"os"
)

type Map struct {
	FilePath string
}

func (o Map) Read(options ReadOptions) (Chunk, error) {
	f, err := os.Open(o.FilePath)
	if err != nil {
		return Chunk{}, err
	}
	defer f.Close()

	raw := make([]byte, options.MaxSize)

	_, err = f.ReadAt(raw, options.Offset)
	if err != nil {
		return Chunk{}, err
	}

	if options.DoneChar != 0 {
		remove := false

		for i := 0; i < len(raw); i++ {
			if rune(raw[i]) == options.DoneChar {
				remove = true
			}

			if remove {
				raw = append(raw[:i], raw[i+1:]...)
			}
		}
	}

	return Chunk{
		Data: raw,
	}, nil
}

func (o Map) Write(options WriteOptions) error {
	return errors.New("Not implemented yet")
}

type ReadOptions struct {
	Offset   int64
	MaxSize  int
	DoneChar rune
}

type WriteOptions struct {
	// TODO.
}

type Chunk struct {
	Data []byte
}

func (o Chunk) String() string {
	return string(o.Data)
}

func (o Chunk) Bytes() []byte {
	return o.Data
}

func NewMap(filePath string) (Map, error) {
	return Map{
		FilePath: filePath,
	}, nil
}