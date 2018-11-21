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
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

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

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		bm.Stop()
		fmt.Println(sig)
		fmt.Println("exiting")
		done <- true
	}()
	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	fmt.Println("awaiting signal")
	<-done

}
