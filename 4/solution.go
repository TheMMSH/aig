package fourth

import (
	"fmt"
	"sync"
	"time"
)

const MaxLenBuffer = 20

type Buffer struct {
	m    sync.RWMutex
	Buff []byte
}

func NewBuffer() Buffer {
	return Buffer{
		Buff: make([]byte, 0),
	}
}

func (b *Buffer) read() string {
	b.m.RLock()
	defer b.m.RUnlock()

	return string(b.Buff)
}

func (b *Buffer) write() {
	b.m.Lock()
	defer b.m.Unlock()

	if len(b.Buff) > MaxLenBuffer {
		b.Buff = make([]byte, 0)
	}
	t, _ := generateRandomSequence(4)
	b.Buff = append(b.Buff, t...)
}

func RunReaderWriters(readersCount, writersCount int, done <-chan bool) {
	buff := NewBuffer()
	for i := 0; i < readersCount; i++ {
		readerIndex := i
		go func() {
			for {
				time.Sleep(1 * time.Millisecond)
				select {
				case <-done:
					return
				default:
					fmt.Printf("reader(%d) reads: %s\n", readerIndex, buff.read())
				}
			}
		}()
	}

	for i := 0; i < writersCount; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Millisecond)
				select {
				case <-done:
					return
				default:
					buff.write()
				}
			}
		}()
	}
}
