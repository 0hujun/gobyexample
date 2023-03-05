package gobyexample

import (
	"fmt"
	"testing"
	"time"
)

// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

func TestTimer(t *testing.T) {
	// Timers represent a single event in the future. You
	// tell the timer how long you want to wait, and it
	// provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used
	// `time.Sleep`. One reason a timer may be useful is
	// that you can cancel the timer before it fires.
	// Here's an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the `timer2` enough time to fire, if it ever
	// was going to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}

func TestTicker(t *testing.T) {

	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `select` builtin on the channel to await the
	// values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
