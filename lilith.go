package lilith

import (
	"fmt"
	"sync"
)

// Lilith represents the main structure for the Lilith system
type Lilith struct {
	mu      sync.RWMutex
	running bool
	// Add more fields as needed
}

// New creates a new instance of Lilith
func New() *Lilith {
	return &Lilith{
		running: false,
	}
}

// Start initializes and starts Lilith
func (l *Lilith) Start() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.running {
		return fmt.Errorf("lilith is already running")
	}

	l.running = true
	return nil
}

// Stop gracefully stops Lilith
func (l *Lilith) Stop() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.running {
		return fmt.Errorf("lilith is not running")
	}

	l.running = false
	return nil
}

// IsRunning returns the current state of Lilith
func (l *Lilith) IsRunning() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.running
}

// Example usage:
/*
func main() {
    lilith := New()
    
    err := lilith.Start()
    if err != nil {
        log.Fatal(err)
    }
    
    // Your Lilith implementation here
    
    err = lilith.Stop()
    if err != nil {
        log.Fatal(err)
    }
}
*/