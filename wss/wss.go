package wss

import (
	"bytes"
	"encoding/json"
	"github.com/godcong/qtrago/util"
	"log"
)

type HandlerFunc func(p util.Map) error

type Responder interface {
	ToMap() util.Map
	Bytes() []byte
}

type Requester interface {
	Request() ([]byte, error)
}

type Reader interface {
	Read() ([]byte, error)
}

type Writer interface {
	Write(p util.Map) error
}

type ReadWrite interface {
	Reader
	Writer
}

type Clientele interface {
	Requester
	Responder
}

type WebSocketNotify interface {
	ReadWrite
	NeedRead() bool
	NeedWrite() util.Map
	Type() string
	CallBack(p util.Map) error
}

type response struct {
	buff *bytes.Buffer
}

func (r *response) ToMap() util.Map {
	maps := util.Map{}
	err := json.Unmarshal(r.Bytes(), &maps)
	if err != nil {
		log.Println(err)
		return nil
	}
	return maps
}

func (r *response) Bytes() []byte {
	return r.buff.Bytes()
}

func ReadToResponder(r Reader) Responder {
	b, err := r.Read()
	if err != nil {
		return nil
	}
	resp := &response{
		buff: bytes.NewBuffer(b),
	}
	return resp
}
