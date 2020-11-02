package semaphore

import "errors"

// Semaphore struct which handles Throttling of Concurrency
type Semaphore struct {
	s chan bool
}

// Acquire semaphore Lock
func (S *Semaphore) Acquire() {
	S.s <- true
}

// Release semaphore Lock
func (S *Semaphore) Release() {
	if len(S.s) > 0 {
		<-S.s
	}
}

// NewSemaphore returns a semaphore of size {size}
func NewSemaphore(size int) (*Semaphore, error) {
	if size < 1 {
		return nil, errors.New("Size must be a positive integer")
	}
	return &Semaphore{s: make(chan bool, size)}, nil
}
