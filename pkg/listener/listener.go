package listener

import (
	"net"

	"github.com/mehmetolgundev/splitter/pkg/parser"
	"github.com/mehmetolgundev/splitter/pkg/splitter"
)

type Listener interface {
	Listen() error
}

type tcpListener struct {
}

func NewTCPListener() Listener {
	return &tcpListener{}
}

func (tl *tcpListener) Listen() error {
	l, err := net.Listen("tcp", ":1453")
	if err != nil {
		return err
	}
	defer l.Close()
	for {
		con, err := l.Accept()
		if err != nil {
			return nil
		}
		defer con.Close()
		//SPLIT filename chunksize
		//Total size is 64 bytes
		msg := make([]byte, 64)

		_, err = con.Read(msg)
		if err != nil {
			return err
		}
		req, err := parser.Parse(string(msg))
		if err != nil {
			con.Write([]byte(err.Error()))
			continue
		}
		c := splitter.Split(req.Filename, req.ChunkSize)
		r := <-c
		if r.Err != nil {
			con.Write([]byte(err.Error()))
			continue
		}
		con.Write([]byte("File is created"))

	}

}
