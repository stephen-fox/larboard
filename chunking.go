package larboard

import (
	"os"
)

type FileChunkStringer struct {
	FilePath string
	Offset   int64
	MaxSize  int
	DoneChar int32
}

func (o FileChunkStringer) Read() (string, error) {
	f, err := os.Open(o.FilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return o.ReadOpen(f)
}

func (o FileChunkStringer) ReadOpen(f *os.File) (string, error) {
	raw := make([]byte, o.MaxSize)

	_, err := f.ReadAt(raw, o.Offset)
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