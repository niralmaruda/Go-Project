package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	name := "Elodie"
	name = updateName(name)
	fmt.Println(name)
}

func updateName(name string) string {
	name = "Niral"
	return name
}
