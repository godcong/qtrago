package wss

import (
	"context"
	"github.com/godcong/qtrago/util"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type BitMexWSS struct {
	context.Context
	cancel context.CancelFunc
	conn   *websocket.Conn
	f      HandlerFunc
	Host   string
}

func (wss *BitMexWSS) Request() ([]byte, error) {

	_, message, err := wss.conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return nil, err
	}
	log.Printf("recv: %s", message)
	return message, nil
}

func (wss *BitMexWSS) Type() string {
	return "bitmex"
}

func (wss *BitMexWSS) Handle(f HandlerFunc) *BitMexWSS {
	wss.f = f
	return wss
}

func (wss *BitMexWSS) CallBack(p util.Map) error {
	if wss.f != nil {
		return wss.f(p)
	}
	return nil
}

func (wss *BitMexWSS) Start() error {
	var err error
	wss.conn, _, err = websocket.DefaultDialer.Dial(wss.Host, nil)
	if err != nil {
		log.Fatal("dial:", err)
		return err
	}
	wss.Context, wss.cancel = context.WithCancel(context.Background())
	err = wss.conn.WriteJSON(util.Map{
		"op":   "subscribe",
		"args": []string{"orderBookL2:XBTUSD"},
	})
	if err != nil {
		log.Println("write:", err)
		return err
	}
	notify(wss.Context, wss)
	return nil
}

func (wss *BitMexWSS) Stop() {
	wss.cancel()
}

func notify(ctx context.Context, notify WebSocketNotify) {
	for {
		select {
		case <-ctx.Done():
			log.Println("end: ", ctx.Err(), notify.Type())
			return
		default:
		}
		log.Println("notify")
		resp := RequesterToResponder(notify)
		if resp != nil {
			notify.CallBack(resp.ToMap())
		}
		time.Sleep(1000)
	}
}
