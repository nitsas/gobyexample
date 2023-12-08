package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ticker *time.Ticker, sigs <-chan os.Signal, done chan<- bool) {
	logger := log.New(os.Stdout, "worker:", log.LstdFlags|log.Lmsgprefix)

	for {
		select {
		case sig := <-sigs:
			logger.Println("Got signal:", sig)
			logger.Println("Exiting...")
			done <- true
			return
		case <-ticker.C:
			logger.Println("Working")
		}
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("main:")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("This program will be notified about SIGINT and SIGTERM and shutdown gracefully")

	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool, 1)

	go worker(ticker, sigs, done)

	log.Println("Waiting for worker... Kill it with Ctrl-C.")
	<-done
	log.Println("Exiting...")
}
