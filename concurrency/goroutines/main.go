package main

import (
	"log"
	"strconv"
	"sync"
)

func AddNumbersToString(from, to int, wg *sync.WaitGroup, mtx *sync.Mutex, res *string) {
	mtx.Lock()
	for i := from; i < to; i++ {
		*res = *res + "|" + strconv.Itoa(i) + "|"
	}
	mtx.Unlock()
	wg.Done()
}

func main() {
	var res string
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}
	wg.Add(4)
	go AddNumbersToString(1, 25, wg, mtx, &res)
	go AddNumbersToString(26, 50, wg, mtx, &res)
	go AddNumbersToString(51, 75, wg, mtx, &res)
	go AddNumbersToString(76, 100, wg, mtx, &res)

	wg.Wait()
	log.Println(res)
}
