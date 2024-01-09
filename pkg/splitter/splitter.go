package splitter

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
func Concat(files [][]byte, filename string) {
	var file []byte
	for _, f := range files {
		file = append(file, f...)
	}
	err := os.WriteFile(filename, file, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
func ConcatFiles(prefixFileName string, extension string, newFileName string) {
	var file []byte
	counter := 0
	var err error
	for err == nil {
		fileName := fmt.Sprintf("%s%d.%s", prefixFileName, counter, extension)
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			break
		}
		b, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		file = append(file, b...)
		counter += 1
	}
	err = os.WriteFile(newFileName, file, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
