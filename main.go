package main

import (
	"fmt"
	"github.com/godcong/qtrago/util"
	"github.com/godcong/qtrago/wss"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var err error
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	/*
		bm := &wss.BitMexWSS{
			Host: "wss://www.bitmex.com/realtime",
		}
		bm.Handle(func(p util.Map) error {
			log.Println("print callback ", p)
			return nil
		})
		err = bm.Start([]string{"trade:XBTUSD", "instrument:XBTUSD", "trade:ETHUSD", "instrument:ETHUSD", "trade:TRXZ18", "instrument:TRXZ18"})
		if err != nil {
			log.Println(err)
			return
		}
	*/
	bm2 := &wss.BitMexWSS{
		Host: "wss://testnet.bitmex.com/realtime",
	}

	bm2.Handle(func(p util.Map) error {
		log.Println("print callback2 ", p)
		return nil
	})
	err = bm2.Start(util.Map{
		"op":   "authKeyExpires",
		"args": "trade:TRXZ18",
	})
	if err != nil {
		log.Println(err)
		return
	}

	q := &qtra{}
	go q.Start()

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		//bm.Stop()
		q.Stop()
		bm2.Stop()
		fmt.Println(sig)
		fmt.Println("exiting")
		done <- true
	}()
	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	//fmt.Println("awaiting signal")
	//time.Sleep(5 * time.Second)
	//, "instrument:XBTUSD", "trade:ETHUSD", "instrument:ETHUSD", "trade:TRXZ18", "instrument:TRXZ18"
	bm2.PutWrite(util.Map{
		"op":   "subscribe",
		"args": "trade:TRXZ18",
	})
	//time.Sleep(5 * time.Second)
	bm2.PutWrite(util.Map{
		"op":   "subscribe",
		"args": "trade:XBTUSD",
	})
	bm2.PutWrite(util.Map{
		"op":   "subscribe",
		"args": "instrument:XBTUSD",
	})
	//time.Sleep(5 * time.Second)
	bm2.PutWrite(util.Map{
		"op":   "subscribe",
		"args": "instrument:ETHUSD",
	})
	//time.Sleep(5 * time.Second)
	bm2.PutWrite(util.Map{
		"op":   "subscribe",
		"args": "instrument:TRXZ18",
	})

	<-done

}
