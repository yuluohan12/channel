package main

import (
	"fmt"
	"sync"
	"time"
)

var w1 sync.WaitGroup
var w2 sync.WaitGroup
var w3 sync.WaitGroup

func main() {
	w1.Add(10)
	w2.Add(5)
	w3.Add(3)

	填装 := make(chan bool, 1)
	瞄准 := make(chan bool)
	开炮 := make(chan bool)

	填装 <- true

	for i := 0; i < 10; i++ { //打印四组，三个goroutine需要执行4次
		go 填装_1(&w1, 填装, 瞄准)

	}
	for j := 0; j < 5; j++ {
		go 瞄准_2(&w2, 瞄准, 开炮)

	}
	for k := 0; k < 3; k++ {
		go 开炮_3(&w3, 开炮, 填装)
	}

	w1.Wait()
	w2.Wait()
	w3.Wait()
}

func 填装_1(w1 *sync.WaitGroup, 填装 chan bool, 瞄准 chan bool) {
	if ok := <-填装; ok {
		fmt.Print("填装->")
		time.Sleep(1 * time.Second)
		瞄准 <- true
	}
}

func 瞄准_2(w2 *sync.WaitGroup, 瞄准 chan bool, 开炮 chan bool) {
	if ok := <-瞄准; ok {
		fmt.Print("瞄准->")
		time.Sleep(1 * time.Second)
		开炮 <- true
	}
}

func 开炮_3(w3 *sync.WaitGroup, 开炮 chan bool, 填装 chan bool) {
	if ok := <-开炮; ok {
		fmt.Println("开炮!")

		填装 <- true
	}
}
