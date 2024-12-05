package main

import (
	"sync"
	"testing"
)

func TestSample(t *testing.T) {
	q := NewQutex()
	q.Lock()
	q.Unlock()
}

func TestLockTwice(t *testing.T) {
	q := NewQutex()
	var wg sync.WaitGroup
	// defer recover()
	wg.Add(2)
	go func() {
		defer wg.Done()
		q.Lock()
	}()
	go func() {
		defer wg.Done()
		q.Lock()
	}()
	wg.Wait()
}
