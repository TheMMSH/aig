package main

import (
	fourth "aig/4"
	"time"
)

const WaitDurationForTest = 10000 * time.Millisecond

func main() {
	done := make(chan bool)
	fourth.RunReaderWriters(8, 2, done)
	time.Sleep(WaitDurationForTest)
	for i := 0; i < 10; i++ {
		done <- true
	}
}
