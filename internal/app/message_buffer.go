package app

import (
	"sync"
)

type messageBuffer struct {
	mu       sync.RWMutex
	size     int
	cursor   int // Points to the first free element.
	messages [][]byte
}

func newMessageBuffer(size int) *messageBuffer {
	return &messageBuffer{
		size:     size,
		messages: make([][]byte, size),
	}
}

func (b *messageBuffer) Add(item []byte) {
	b.mu.Lock()
	b.messages[b.cursor] = item
	b.cursor = (b.cursor + 1) % b.size
	b.mu.Unlock()
}

func (b *messageBuffer) Iterate(fn func([]byte) error) error {
	b.mu.RLock()

	for i := b.cursor + 1; i < len(b.messages); i++ {
		if err := fn(b.messages[i]); err != nil {
			return err
		}
	}

	for i := 0; i <= b.cursor; i++ {
		if err := fn(b.messages[i]); err != nil {
			return err
		}
	}

	b.mu.RUnlock()

	return nil
}
