package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// 4. Create a function to read a file and pipe each line to a channel. Be careful here,
// though; you may need to add a WaitGroup or something else to avoid any
// deadlocks.
func source(filepath string, out chan int, wg *sync.WaitGroup) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rl := bufio.NewReader(file)
	for {
		line, err := rl.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				wg.Done()
				return
			} else {
				panic(err)
			}
		}
		rline := strings.ReplaceAll(line, "\n", "")
		rline = strings.ReplaceAll(rline, "\r", "")
		digit, err := strconv.Atoi(rline)
		if err != nil {
			panic(err)
		}
		out <- digit
	}
}

// 5. Create a function to receive the numbers and pipe the odd numbers to one channel
// and the even numbers to another channel.
func split(in, outOdd, outEven chan int, wg *sync.WaitGroup) {
	for digit := range in {
		if digit%2 == 0 {
			outEven <- digit
		} else {
			outOdd <- digit
		}
	}
	close(outEven)
	close(outOdd)
	wg.Done()
}

// 6. Create a function to sum the numbers and pipe the result to a new channel.
func sum(in, out chan int, wg *sync.WaitGroup) {
	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
	wg.Done()
}

// 7. Create a merging function to read from the odd and even channels and write to a
// file called result.txt. Each line in this file should contain the word "Odd" or "Even,"
// depending on the value, followed by the sum.
func merge(inOdd, inEven chan int, wg *sync.WaitGroup) {
	resFile, err := os.OpenFile("result.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer resFile.Close()

	for i := 0; i < 2; i++ {
		select {
		case i := <-inOdd:
			resFile.Write([]byte(fmt.Sprintf("Odd %d\n", i)))
		case i := <-inEven:
			resFile.Write([]byte(fmt.Sprintf("Even %d\n", i)))
		}
	}
	wg.Done()
}

// 8. Create the main function to run all the Goroutines and handle the WaitGroups, if
// needed.
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	wg2 := &sync.WaitGroup{}
	wg2.Add(4)

	digits := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenChSum := make(chan int)
	oddChSum := make(chan int)

	go source("text1.txt", digits, wg)
	go source("text2.txt", digits, wg)
	go split(digits, oddCh, evenCh, wg2)
	go sum(oddCh, oddChSum, wg2)
	go sum(evenCh, evenChSum, wg2)
	go merge(oddChSum, evenChSum, wg2)

	wg.Wait()
	close(digits)
	wg2.Wait()
	close(evenChSum)
	close(oddChSum)
}
