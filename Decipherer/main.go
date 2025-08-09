package main

import (
	"sync"
)

func StartDecipher(senderChan chan string, decipherer func(encrypted string) string) chan string {
	mu := sync.Mutex{}
	deciphered := make(chan string, 5)

	go func() {
		for encrypted := range senderChan {	
			mu.Lock()
			deciphered <- decipherer(encrypted)
			mu.Unlock()
		}
	}()
	return deciphered
}
