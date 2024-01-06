package main

import (
	"bufio"
	"io"
	"os"
)

type Result struct {
	Err   error
	Files [][]byte
}

func Split(filePath string, chunkSize int) <-chan Result {
	c := make(chan Result)
	go func() {
		result := Result{}
		f, err := os.Open(filePath)
		if err != nil {
			c <- Result{
				Err: err,
			}
			return
		}
		r := bufio.NewReader(f)
		for {
			chunk := make([]byte, chunkSize)

			n, err := r.Read(chunk)

			if err != nil {
				if err == io.EOF {
					break
				}
				c <- Result{
					Err: err,
				}
				return
			}

			result.Files = append(result.Files, chunk[:n])
		}
		c <- result
	}()

	return c
}
