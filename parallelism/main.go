package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go imprimirNumeroParalelamente(i, &wg)
	}

	// Espera até que todas as goroutines terminem
	wg.Wait()
}

func imprimirNumeroParalelamente(numero int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Número:", numero)
}

