package main

import (
	"github.com/godcong/qtrago/util"
	"github.com/godcong/qtrago/wss"
	"log"
	"time"
)

func main() {
	bm := &wss.BitMexWSS{
		Host: "wss://www.bitmex.com/realtime",
	}
	bm.Handle(func(p util.Map) error {
		log.Println("print callback ", p)
		return nil
	})
	err := bm.Start()
	if err != nil {
		log.Println(err)
		return
	}
	time.Sleep(time.Second * 100)
	bm.Stop()
	time.Sleep(time.Second * 10)
}
