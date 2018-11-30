package wss

import (
	"context"
	"github.com/godcong/qtrago/util"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type BitMexWSS struct {
	context.Context
	cancel context.CancelFunc
	conn   *websocket.Conn
	f      HandlerFunc
	Host   string
	pool   sync.Pool
	//needWrite bool
	//writeData util.Map
}

func (b *BitMexWSS) Pool() sync.Pool {
	return b.pool
}

func (b *BitMexWSS) PutWrite(p util.Map) {
	b.pool.Put(p)
}

func (b *BitMexWSS) NeedRead() bool {
	return true
}

func (b *BitMexWSS) NeedWrite() util.Map {
	v := b.pool.Get()
	if v != nil {
		return v.(util.Map)
	}
	return nil
}

func (b *BitMexWSS) Read() ([]byte, error) {
	_, message, err := b.conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return nil, err
	}
	log.Printf("recv: %s", message)
	return message, nil
}

func (b *BitMexWSS) Write(json util.Map) error {
	log.Println("write json", json)
	return b.conn.WriteJSON(json)
}

func (b *BitMexWSS) Type() string {
	return "bitmex"
}

func (b *BitMexWSS) Handle(f HandlerFunc) *BitMexWSS {
	b.f = f
	return b
}

func (b *BitMexWSS) WriteJSON(p util.Map) error {
	//json := p.ToJSON()
	//log.Println("write json", string(json))
	return b.conn.WriteJSON(p)
}

func (b *BitMexWSS) CallBack(p util.Map) error {
	if b.f != nil {
		return b.f(p)
	}
	return nil
}

func (b *BitMexWSS) Start(op string, list []string) error {
	var err error
	b.conn, _, err = websocket.DefaultDialer.Dial(b.Host, nil)
	if err != nil {
		log.Fatal("dial:", err)
		return err
	}
	b.Context, b.cancel = context.WithCancel(context.Background())
	//first send subscribe
	err = b.WriteJSON(util.Map{
		"op":   op,
		"args": list,
	})
	if err != nil {
		log.Println("write:", err)
		return err
	}
	go notify(b.Context, b)
	return nil
}

func (b *BitMexWSS) Stop() {
	b.cancel()
}

func notify(ctx context.Context, notify WebSocketNotify) {
	for {
		log.Println("notify")
		select {
		case <-ctx.Done():
			log.Println("end: ", ctx.Err(), notify.Type())
			return
		default:
		}
		if data := notify.NeedWrite(); data != nil {
			_ = notify.Write(data)
		}

		if notify.NeedRead() {
			resp := ReadToResponder(notify)
			if resp != nil {
				r := resp.ToMap()
				err := notify.CallBack(r)
				if err != nil {
					log.Println("callback err", err)
				}
			}
		}

		time.Sleep(time.Duration(500) * time.Nanosecond)
	}
}
