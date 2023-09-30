package main

import (
	"fmt"
	"time"
)

func clock(t *time.Ticker, done chan bool) {
	for i := 1; true; i++ {
		select {
		case <-done:
			// they told us to stop and exit
			fmt.Println("clock: done!")
			time.Sleep(200 * time.Millisecond)
			// signal back that we are ok to exit
			done <- true
			return
		case <-t.C:
			fmt.Println("clock: tick", i)
		}
	}
}

func main() {
	t := time.NewTicker(time.Second)
	done := make(chan bool)

	go clock(t, done)

	fmt.Println("main: sleeping for 4 secs")
	time.Sleep(4 * time.Second)
	fmt.Println("main: woke up! kill the ticker")
	t.Stop()
	fmt.Println("main: kill the clock")
	done <- true

	fmt.Println("main: wait for the clock to tell us it's exiting")
	<-done
	fmt.Println("main: done! bye bye")
}
