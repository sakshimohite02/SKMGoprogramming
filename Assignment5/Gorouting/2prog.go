package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers()
	go printLetters()

time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
}