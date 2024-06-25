package main

import ("fmt")

func loop() {

	// Loops

	i := 1
	num := 10
	for i <= num {

		fmt.Print(i, " ")

		i += 1
		
	}

	fmt.Println("")

	for j := 1; j <= 20; j++ {
		fmt.Print(j, " ")
	}

	fmt.Println("")

	fmt.Println("Range Keyword")

	// range keyword

	for i := range 3 {
		fmt.Print(i, " ")
	}

	fmt.Println("")
	fmt.Println("Break")
	for i := range 4 {
		if i == 3 {
			break
		}
		fmt.Print(i, " ")
	}


	fmt.Println("")
	fmt.Println("Continue")
	for i := range 4 {
		if i == 0 {continue}

		fmt.Print(i, " ")
	}

	for {
		fmt.Println("Manoj")
		break;
	}

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

	if num % 2 == 0 && num > 0 {
		fmt.Println("Divisible by 0 and positive")
	}
}