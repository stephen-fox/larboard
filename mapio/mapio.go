package mapio

import (
	"bufio"
	"errors"
	"os"
)

type Mapper struct {
	FilePath string
}

func (o Mapper) Read(options ReadOptions) (Chunk, error) {
	f, err := os.Open(o.FilePath)
	if err != nil {
		return Chunk{}, err
	}
	defer f.Close()

	_, err = f.Seek(options.Offset, 0)
	if err != nil {
		return Chunk{}, err
	}

	r := bufio.NewReaderSize(f, options.MaxSize)
	var raw []byte

	if options.Stop {
		raw, err = r.ReadBytes(byte(options.StopChar))
		if err != nil {
			return Chunk{}, err
		}

		if len(raw) > 0 {
			raw = raw[:len(raw)-1]
		}
	} else {
		raw = make([]byte, options.MaxSize)

		_, err = r.Read(raw)
		if err != nil {
			return Chunk{}, err
		}
	}

	return Chunk{
		Data: raw,
	}, nil
}

func (o Mapper) Write(options WriteOptions) error {
	return errors.New("Not implemented yet")
}

type ReadOptions struct {
	Offset   int64
	MaxSize  int
	Stop     bool
	StopChar rune
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

func NewMap(filePath string) (Mapper, error) {
	return Mapper{
		FilePath: filePath,
	}, nil
}