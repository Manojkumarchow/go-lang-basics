package main

import ("fmt";"math")

const surname string = "Pentela"

func variables() {
	fmt.Println("Hello, World!")

	var firstName = "Manoj"

	fmt.Println("First Name: ", firstName)

	var a, b int = 0, 0
	fmt.Println("a, b: ", a, b)

	lastName := "Kumar"

	fmt.Println("Last Name: ", lastName)

	fmt.Println("Surname: ", surname)

	const n = 100

	fmt.Println("Sin(n): ", math.Sin(n));
}