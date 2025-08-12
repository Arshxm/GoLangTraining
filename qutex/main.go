package main

// DO NOT USE ANY IMPORT

type Qutex struct {
	ch chan struct{}
}

func NewQutex() *Qutex {
	return &Qutex{
		ch: make(chan struct{}, 1),
	}
}

func (qu *Qutex) Lock() {
	qu.ch <- struct{}{}
}

func (qu *Qutex) Unlock() {
	select {
	case <-qu.ch:
	default:
		panic("unlock of unlocked mutex")
	}
}
