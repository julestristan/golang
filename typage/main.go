package main

import "fmt"

func main() {
	WriteMessage("Test\n")
	fmt.Printf(AppendMessage("Hello", "World"))
}

func WriteMessage(message string) {
	fmt.Printf("voici le message: %v", message)
}

func AppendMessage(m1 string, m2 string) string {
	return fmt.Sprintf("Message 1: %v\nMessage 2: %v", m1, m2)
}
