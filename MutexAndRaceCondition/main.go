package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Mutex and Race conditions")
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(5)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		fmt.Println("One Routine")
		score = append(score, 1)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Third Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Fourth Routine")
		mut.Lock()
		score = append(score, 4)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	go func(mut *sync.RWMutex) {
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		defer wg.Done()
	}(mut)
	wg.Wait()
	fmt.Println(score)
}
