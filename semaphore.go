package semaphore

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
	<-S.s
}

// NewSemaphore returns a semaphore of size {size}
func NewSemaphore(size int) *Semaphore {
	return &Semaphore{s: make(chan bool, size)}
}
