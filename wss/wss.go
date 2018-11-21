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

type Clientele interface {
	Requester
	Responder
}

type WebSocketNotify interface {
	Requester
	Type() string
	CallBack(p util.Map) error
}

type response struct {
	buff *bytes.Buffer
}

func (r *response) ToMap() util.Map {
	maps := util.Map{}
	err := json.Unmarshal(r.buff.Bytes(), &maps)
	if err != nil {
		log.Println(err)
		return nil
	}
	return maps
}

func (*response) Bytes() []byte {
	panic("implement me")
}

func RequesterToResponder(r Requester) Responder {
	b, err := r.Request()
	if err != nil {
		return nil
	}
	resp := &response{
		buff: bytes.NewBuffer(b),
	}
	return resp
}
