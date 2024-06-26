package main

import (
	"fmt"
)

// mencetak bilangan prima yang kurang dari 10 rb

func first_task() {
	limit := 10_000

	for num := 2; num < limit; num++ {
		bilanganPrima := true

		for i := 2; i * i <= num;i++ {
			if num % i == 0 {
				bilanganPrima = false
				break
			}
		}

		if bilanganPrima {
			fmt.Println(num)
		}
	}
}