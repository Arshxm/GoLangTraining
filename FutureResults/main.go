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
		fResult.ResultChan <- result
		fResult.Done.Store(true)
	}()
	return fResult
}
func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
	fResult := &FutureResult{
		ResultChan: make(chan string),
	}
	go func() {
		// Create a channel to receive the task result
		taskChan := make(chan string, 1)

		// Start the task in a separate goroutine and measure its duration
		go func() {
			start := time.Now()
			result := t()
			fResult.Duration = time.Since(start)
			taskChan <- result
		}()

		// Wait for either the task to complete or timeout
		select {
		case result := <-taskChan:
			fResult.ResultChan <- result
		case <-time.After(timeout):
			fResult.ResultChan <- "timeout"
		}
		fResult.Done.Store(true)
	}()
	return fResult
}

func (fResult *FutureResult) Await() string {
	select {
	case result := <-fResult.ResultChan:
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
		}
	}()
	return combinedResult
}
