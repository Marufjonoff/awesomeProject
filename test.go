package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	print(Hello("Obidjondan salom"))
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
