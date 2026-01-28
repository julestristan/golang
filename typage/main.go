package main

import "fmt"

func main() {
	WriteMessage("Test")
}
func WriteMessage(message string) {
	fmt.Printf("voici le message: %v", message)
}
