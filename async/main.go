package main

import (
	"sync/atomic"
	"time"
)

type FutureResult struct {
	Done       atomic.Bool
	ResultChan chan string
}

func NewFutureResult(chanSize ...int) *FutureResult {
	if len(chanSize) == 0 {
		chanSize = append(chanSize, 1)
	} else if chanSize[0] <= 0 {
		chanSize[0] = 1
	}
	result := FutureResult{
		atomic.Bool{},
		make(chan string, chanSize[0]),
	}
	result.Done.Store(false)
	return &result
}

type Task func() string

func Async(t Task) *FutureResult {
	var result = NewFutureResult()
	go func() {
		result.ResultChan <- t()
		result.Done.Store(true)
	}()
	return result
}

func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
	fResult := NewFutureResult()
	resultChan := make(chan string, 1)

	go func() {
		result := t()
		resultChan <- result
	}()

	go func() {
		select {
		case res := <-resultChan:
			fResult.ResultChan <- res
			fResult.Done.Store(true)
		case <-time.After(timeout):
			fResult.ResultChan <- "timeout"
			fResult.Done.Store(true)
		}
	}()

	return fResult
}

func (fResult *FutureResult) Await() string {
	select {
	case res := <-fResult.ResultChan:
		{
			fResult.Done.Store(true)
			return res
		}
	}
}

func CombineFutureResults(fResults ...*FutureResult) *FutureResult {
	result := NewFutureResult(len(fResults))
	result.Done.Store(true)
	for _, item := range fResults {
		result.ResultChan <- <-item.ResultChan
	}

	return result
}
