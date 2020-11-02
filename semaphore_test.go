package semaphore

import "testing"
import "fmt"

func TestItFillsUpAsExpected(t *testing.T) {
	size := 1
	s, _ := NewSemaphore(size)
	s.Acquire()
	if len(s.s) != size {
		t.Error()
	}
}

func TestItReleasesProperly(t *testing.T) {
	size := 2
	s, _ := NewSemaphore(size)
	s.Acquire() // +1
	s.Acquire() // +1
	s.Release() // -1
	if len(s.s) != 1 {
		t.Error()
	}
}

func TestItOnlyAllowsPositiveIntegerSize(t *testing.T) {
	notValidSize := -1
	validSize := 1
	zero := 0
	_, err := NewSemaphore(notValidSize)
	if err == nil {
		t.Error()
	}
	_, err = NewSemaphore(zero)
	if err == nil {
		t.Error()
	}
	_, err = NewSemaphore(validSize)
	if err != nil {
		t.Error()
	}
}

func TestItDoesntHangOnEmptyRelease(t *testing.T) {
	s, _ := NewSemaphore(1)
	fmt.Printf("Size: %d\n", len(s.s))
	s.Release()
}
