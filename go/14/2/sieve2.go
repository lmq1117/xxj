package main

import "fmt"

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}

func generateCh() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filterInt(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

//返回包含素数的channel
func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generateCh()
		for {
			prime := <-ch
			ch = filterInt(ch, prime)
			out <- prime
		}
	}()
	return out
}
