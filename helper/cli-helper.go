package helper

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ShowAuthMenu() {
	fmt.Println("=== REA Game Store Authentication CLI ===")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println("=====================")
}

func ShowMainMenu() {
	fmt.Println("\n=== REA Game Store Main Menu ===")
	fmt.Println("1. Game Management")
	fmt.Println("2. Order Management")
	fmt.Println("3. Reports")
	fmt.Println("4. Logout")
	fmt.Println("5. Exit")
	fmt.Println("==========================")
}

func HandleGameOperations(){
	
}
