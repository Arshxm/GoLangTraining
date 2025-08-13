package main

import (
	"sync/atomic"
	"time"
)

type FutureResult struct {
	Done       atomic.Bool
	ResultChan chan string
	Duration   time.Duration // Track how long the task took to execute
}

type Task func() string

func Async(t Task) *FutureResult {
	fResult := &FutureResult{
		ResultChan: make(chan string),
	}
	go func() {
		start := time.Now()
		result := t()
		fResult.Duration = time.Since(start)
		fResult.Done.Store(true)
		fResult.ResultChan <- result
	}()
	return fResult
}
func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
	fResult := &FutureResult{
		ResultChan: make(chan string),
	}
	go func() {
		taskChan := make(chan string, 1)
		go func() {
			start := time.Now()
			result := t()
			fResult.Duration = time.Since(start)
			taskChan <- result
		}()
		select {
		case result := <-taskChan:
			fResult.ResultChan <- result
			fResult.Done.Store(true)
		case <-time.After(timeout):
			fResult.ResultChan <- "timeout"
		}
	}()
	return fResult
}

func (fResult *FutureResult) Await() string {
	select {
	case result := <-fResult.ResultChan:
		fResult.Done.Store(true)
		return result
	case <-time.After(10 * time.Second):
		return "timeout"
	}
}

func CombineFutureResults(fResults ...*FutureResult) *FutureResult {
	combinedResult := &FutureResult{
		ResultChan: make(chan string),
	}
	go func() {
		for _, fResult := range fResults {
			combinedResult.ResultChan <- <-fResult.ResultChan
			fResult.Done.Store(true)
		}
	}()
	return combinedResult
}
