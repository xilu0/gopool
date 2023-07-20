package gopool

import (
	"fmt"
	"testing"
)

func TestNewPoolFactory(t *testing.T) {
	pf := NewPoolFactory()
	if _, ok := pf.(*poolFactory); !ok {
		t.Errorf("NewPoolFactory() did not return a poolFactory instance")
	}
}

func TestCreatePool(t *testing.T) {
	pf := NewPoolFactory()
	size := 10
	pool := pf.CreatePool(size)
	if pool == nil {
		t.Errorf("Expected a non-nil pool, but got nil")
	}
}

func TestPoolFactory(t *testing.T) {
	factory := NewPoolFactory()

	pool := factory.CreatePool(10)
	go pool.Run()
	defer pool.Stop()
	// Test that pool executes tasks concurrently
	numTasks := 100
	taskFunc := func() error {
		fmt.Println("do taskFunc")
		return nil
	}
	for i := 0; i < numTasks; i++ {
		pool.Add(taskFunc)
	}
}
