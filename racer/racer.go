package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
func Racer(a string, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
*/
func Racer(a, b string) (winner string, err error) {
	/*timeout := make(chan int)
	go func(timeout chan int) {
		time.Sleep(10 * time.Millisecond)
		timeout <- 1
		return
	}(timeout)*/
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
