package splitter

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Result struct {
	Err   error
	Files [][]byte
}

func Split(filePath string, chunkSize byte) <-chan Result {
	c := make(chan Result)
	go func() {
		result := Result{}
		f, err := os.Open(filePath)
		if err != nil {
			c <- Result{
				Err: err,
			}
		}
		defer log.Fatal(f.Close())

		r := bufio.NewReader(f)
		for {
			chunk := make([]byte, chunkSize)
			_, err = r.Read(chunk)
			if err == io.EOF {
				break
			}
			result.Files = append(result.Files, chunk)
		}
		c <- result
	}()

	return c
}
