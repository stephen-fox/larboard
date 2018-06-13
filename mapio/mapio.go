package mapio

import (
	"bufio"
	"io"
	"os"
)

type OpenOptions struct {
	AllowWrite bool
}

type ReadOptions struct {
	Offset int64
	Size   int
	Until  bool
	UntilR rune
}

type WriteOptions struct {
	Offset int64
	Data   []byte
}

type Chunk struct {
	Eof  bool
	Got  int
	Data []byte
}

func (o Chunk) String() string {
	return string(o.Data)
}

func (o Chunk) Bytes() []byte {
	return o.Data
}

func Open(filePath string, options OpenOptions) (*os.File, error) {
	flag := os.O_RDONLY

	if options.AllowWrite {
		flag = os.O_RDWR
	}

	return os.OpenFile(filePath, flag, 0600)
}

func AtomicRead(filePath string, rO ReadOptions) (Chunk, error) {
	f, err := Open(filePath, OpenOptions{})
	if err != nil {
		return Chunk{}, err
	}
	defer f.Close()

	return Read(f, rO)
}

func Read(f *os.File, options ReadOptions) (Chunk, error) {
	_, err := f.Seek(options.Offset, 0)
	if err != nil {
		return Chunk{}, err
	}

	r := bufio.NewReaderSize(f, options.Size)
	var chunk Chunk

	if options.Until {
		raw, err := r.ReadBytes(byte(options.UntilR))
		if err != nil {
			return Chunk{}, err
		}

		if len(raw) > 0 {
			raw = raw[:len(raw)-1]
		}

		chunk.Data = raw
	} else {
		chunk.Data = make([]byte, options.Size)
		var err error

		chunk.Got, err = r.Read(chunk.Data)
		switch err {
		case nil:
			break
		case io.EOF:
			chunk.Eof = true
		default:
			return chunk, err
		}
	}

	return chunk, nil
}

func Write(f *os.File, options WriteOptions) error {
	_, err := f.Seek(options.Offset, 0)
	if err != nil {
		return err
	}

	_, err = f.Write(options.Data)
	if err != nil {
		return err
	}

	return nil
}
