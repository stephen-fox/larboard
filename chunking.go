package larboard

import (
	"os"
)

type FileChunk struct {
	Offset   int64
	MaxSize  int
	DoneChar int32
}

func (o FileChunk) String(f *os.File) (string, error) {
	raw, err := o.Bytes(f)
	if err != nil {
		return "", err
	}

	var chars []int32

	for _, r := range raw {
		c := int32(r)

		if c == o.DoneChar {
			break
		}

		chars = append(chars, c)
	}

	return string(chars), nil
}

func (o FileChunk) Bytes(f *os.File) ([]byte, error) {
	raw := make([]byte, o.MaxSize)

	_, err := f.ReadAt(raw, o.Offset)
	if err != nil {
		return raw, err
	}

	return raw, nil
}