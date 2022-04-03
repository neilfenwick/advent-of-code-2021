package files

import (
	"io"
	"os"
)

func ReadInputStream() (io.ReadCloser, error) {
	var (
		file *os.File
		err  error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			return nil, err
		}
	}
	return file, nil
}
