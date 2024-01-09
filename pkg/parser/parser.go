package parser

import (
	"errors"
	"strconv"
	"strings"
)

type Request struct {
	Filename  string
	ChunkSize int
}

func Parse(msg string) (*Request, error) {
	var request *Request
	v := strings.Split(msg, " ")

	if len(v) != 3 {
		return nil, errors.New("Message's lenght is invalid")
	}

	if v[0] != "SPLIT" {
		return nil, errors.New("Message invalid.")
	}
	chunkSize, err := strconv.Atoi(v[1])
	if err != nil {
		return nil, errors.New("Chunk size invalid")
	}

	request = &Request{
		Filename:  v[1],
		ChunkSize: chunkSize,
	}
	return request, nil

}
