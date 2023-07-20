// BEGIN: 8b9f7e7d8c8b
package gopool

import (
	"errors"
	"sync"
	"testing"
)

func TestPool_Run(t *testing.T) {
	p := NewPool(2)
	go p.Run()
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		p.Add(func() error {
			wg.Done()
			return nil
		})
	}
	wg.Add(1)
	p.Add(func() error {
		wg.Done()
		return errors.New("test error")
	})
	wg.Wait()
	p.Stop()
}

// END: 8b9f7e7d8c8b
