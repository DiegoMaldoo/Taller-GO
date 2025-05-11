package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	if n < 0 {
		return 0
	}
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
		time.Sleep(50 * time.Millisecond)
	}
	return resultado
}

func calcularFactorial(n int, ch chan<- int) {
	res := factorial(n)
	ch <- res
}

func main() {
	numeros := []int{5, 7, 8, 6, 19}

	resultCh := make(chan int, len(numeros))

	for _, n := range numeros {
		go calcularFactorial(n, resultCh)
	}

	for i := 0; i < len(numeros); i++ {
		resultado := <-resultCh
		fmt.Printf("Factorial calculado: %d\n", resultado)
	}
}
