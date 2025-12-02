package main

import (
	"fmt"
)

func clearScreen() {
    fmt.Print("\033[H\033[2J") // Clear screen for better UX
}

func main(){
	clearScreen()
	fmt.Println("Starting REA Game Store Cashier App")
}