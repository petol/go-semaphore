# go-semaphore
Simple Semaphore Solution based on a buffered Channel, Blocking until a semaphore is released


## Usage:
```go
  package main
  import "github.com/petol/go-semaphore"
  
  func main() {
    size_of_semaphore := 3
    s := semaphore.newSemaphore(size_of_semaphore) // Create a new Semaphore of Size 3
    s.Acquire() // Acquire one Semaphore, Available = 3 - 1
    s.Release() // Releases Semaphore, Available = 3
  }
```
