package fourth

import (
	"testing"
	"time"
)

const WaitDurationForTest = 100 * time.Millisecond

func TestForAllTheM8N2(t *testing.T) {
	done := make(chan bool)
	RunReaderWriters(8, 2, done)
	time.Sleep(WaitDurationForTest)
	for i := 0; i < 10; i++ {
		done <- true
	}
}

func TestForAllTheM8N8(t *testing.T) {
	done := make(chan bool)
	RunReaderWriters(8, 8, done)
	time.Sleep(WaitDurationForTest)
	for i := 0; i < 16; i++ {
		done <- true
	}
}

func TestForAllTheM8N16(t *testing.T) {
	done := make(chan bool)
	RunReaderWriters(8, 16, done)
	time.Sleep(WaitDurationForTest)
	for i := 0; i < 24; i++ {
		done <- true
	}
}

func TestForAllTheM2N8(t *testing.T) {
	done := make(chan bool)
	RunReaderWriters(2, 8, done)
	time.Sleep(WaitDurationForTest)
	for i := 0; i < 10; i++ {
		done <- true
	}
}
