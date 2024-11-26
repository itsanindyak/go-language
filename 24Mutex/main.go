package main

import (
	"fmt"
	"sync"
)

type waitGroup = sync.WaitGroup
type mutex = sync.Mutex

func main() {

	wg := &waitGroup{}
	mut := &mutex{}
	var score = []int{0}

	wg.Add(3)
	go func(wg *waitGroup, m *mutex) {
		defer wg.Done()
		m.Lock()
		fmt.Println("Run go routine : 1")
		score = append(score, 1)
		m.Unlock()
	}(wg, mut)
	go func(wg *waitGroup, m *mutex) {
		defer wg.Done()
		fmt.Println("Run go routine : 2")
		score = append(score, 2)
	}(wg, mut)
	go func(wg *waitGroup, m *mutex) {
		defer wg.Done()
		fmt.Println("Run go routine : 3")
		score = append(score, 3)
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
