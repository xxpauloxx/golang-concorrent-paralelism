package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		go imprimirNumero(i)
	}

	// Espera um tempo para dar tempo às goroutines de terminar
	time.Sleep(time.Second)
}

func imprimirNumero(numero int) {
	fmt.Println("Número:", numero)
}

