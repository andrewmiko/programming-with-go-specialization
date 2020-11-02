package main

import (
	"fmt"
	"sync"
	"time"
)

const maxEatingCount = 3
const chopsticksCount = 5
const philosophersCount = 5

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number          int
	leftCS, rightCS *Chopstick
}

func (p *Philosopher) Eat(wg *sync.WaitGroup, host chan bool) {
	defer wg.Done()

	for i := 0; i < maxEatingCount; i++ {
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %d\n", p.number)

		time.Sleep(500 * time.Millisecond)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		fmt.Printf("finishing eating %d\n", p.number)

		<-host
	}
}

func main() {
	chopsticks := make([]*Chopstick, chopsticksCount)
	for i := 0; i < chopsticksCount; i++ {
		chopsticks[i] = &Chopstick{}
	}

	philosophers := make([]*Philosopher, philosophersCount)
	for i := 0; i < philosophersCount; i++ {
		philosophers[i] = &Philosopher{i, chopsticks[i], chopsticks[(i+1)%5]}
	}

	wg := new(sync.WaitGroup)
	host := make(chan bool, 2)
	for _, p := range philosophers {
		wg.Add(1)
		go p.Eat(wg, host)
	}

	wg.Wait()
}
