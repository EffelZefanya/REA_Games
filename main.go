package main

import (
	"fmt"
	"rea_games/cli"
	"rea_games/config"
	"rea_games/handler"
	"rea_games/helper"
)

func main() {
	cli.ClearScreen()
	fmt.Println("Starting REA Game Store Cashier App ⏳")

	inputter := helper.NewInputter()

	db := config.ConnectDB()
	defer db.Close()

	authHandler := handler.NewAuthHandler()
	var currentUserId int
	var loggedIn bool

	fmt.Println("🎮 Welcome to REA Game Store Cashier App")
	fmt.Println("================================")

	for {
		if !loggedIn {
			cli.ShowAuthMenu()
			choice := inputter.ReadInt("Choose option: ")

			switch choice {
			case 1:
				cli.ClearScreen()
				fmt.Println("=== User Registration ===")
				userID, err := authHandler.Register()
				if err != nil {
					fmt.Printf("❌ Error: %v\n\n", err)
				} else {
					currentUserId = userID
					loggedIn = true
					fmt.Printf("✅ Welcome! Registration successful. User ID: %d\n", userID)
				}
			case 2:
				cli.ClearScreen()
				fmt.Println("=== User Login ===")
				userID, err := authHandler.Login()
				if err != nil {
					fmt.Printf("❌ Error: %v\n", err)
				} else {
					currentUserId = userID
					loggedIn = true
					fmt.Printf("✅ Login successful! Welcome back. User ID: %d\n", userID)
				}
			case 3:
				cli.ClearScreen()
				fmt.Println("👋 Thank you for using REA Game Store Cashier. Goodbye!")
				return
			default:
				fmt.Println("❌ Invalid choice! Please try again.")
			}
		} else {
			fmt.Println("You are now logged in.")
			fmt.Printf("Currently assigned as User ID: %d\n", currentUserId)
			cli.ShowMainMenu()
			choice := inputter.ReadInt("Choose option: ")

			switch choice {
			case 1:
				cli.ClearScreen()
				cli.HandleGameOperations()
			case 2:
				cli.ClearScreen()
				cli.HandleOrderOperations(currentUserId)
			case 3:
				cli.ClearScreen()
				// handleReportOperations()
			case 4:
				cli.ClearScreen()
				fmt.Println("✅ Logged out successfully!")
				loggedIn = false
				currentUserId = 0
			case 5:
				cli.ClearScreen()
				fmt.Println("👋 Thank you for using REA Game Store Cashier App. Goodbye!")
				return
			default:
				fmt.Println("❌ Invalid choice! Please try again.")
			}
		}
	}
}
