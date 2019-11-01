package main

import (
	"fmt"
	"sync"
)

var (

myres = make(map[int]int, 20)

mu    sync.Mutex
)

func factorial(n int) {

	var res = 1

	for i := 1; i <= n; i++ {

		res *= i

	}

	mu.Lock()

	myres[n] = res

	mu.Unlock()
}



func main() {

	for i := 1; i <= 20; i++ {

		go factorial(i)

	}

	mu.Lock()

	for i, v := range myres {

		fmt.Printf("myres[%d] = %d\n", i, v)

	}

	mu.Unlock()

}