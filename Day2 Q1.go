package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(letter string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[letter]++
}

func main() {
	words := []string{"fox", "part", "brown", "harshil"}
	var wordsChan = make(chan string, 5)
	c := Container{counters: make(map[string]int)}
	for i := range words {
		wordsChan <- words[i]
	}
	var wg, wg1 sync.WaitGroup

	doIncrement := func(letter string) {
		c.inc(letter)
		wg.Done()
	}

	callWord := func(word string) {
		wg.Add(len(word))
		for j := 0; j < len(word); j++ {
			go doIncrement(word[j : j+1])
		}
		wg1.Done()
	}

	for len(wordsChan) > 0 {
		wg1.Add(1)
		go callWord(<-wordsChan)
	}
	wg1.Wait()
	wg.Wait()
	fmt.Println(c.counters)
}
