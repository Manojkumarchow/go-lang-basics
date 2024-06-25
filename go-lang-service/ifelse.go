package main

import ("fmt")

func ifElse() {

	fmt.Println("If-else")

	if 4 % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 100; num > 300 {
		fmt.Println("Number greater than 300")
	} else if num >= 100 {
		fmt.Println("Number equals or greater than 100")
	} else {
		fmt.Println("Number less than 100")
	}

	if num := 4; num % 2 == 0 && num > 0 {
		fmt.Println("Divisible by 0 and positive")
	}
}