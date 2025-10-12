package main

import (
	"fmt"
	"sync"
	"time"
)

func sig(after time.Duration) chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		time.Sleep(after)
	}()

	return ch
}

func or(channels ...<-chan interface{}) chan interface{} {
	out := make(chan interface{})

	once := sync.Once{}

	for _, ch := range channels {
		go func() {
			select {
			case <-ch:
				once.Do(func() { close(out) })
			case <-out:
			}
		}()

	}

	return out
}

func main() {
	start := time.Now()

	ch1 := sig(time.Hour)
	ch2 := sig(time.Hour * 100)
	ch3 := sig(time.Minute)
	ch4 := sig(time.Second * 3)
	ch5 := sig(time.Minute * 2)

	<-or(ch1, ch2, ch3, ch4, ch5)

	fmt.Println(time.Since(start))
}
