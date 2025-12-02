package main

import (
	"fmt"
	"rea_games/config"
	"rea_games/handler"
	"rea_games/helper"
)

func main(){
	helper.ClearScreen()
	fmt.Println("Starting REA Game Store Cashier App")

	inputter := helper.NewInputter()

	db := config.ConnectDB()
	defer db.Close()

	authHandler := handler.NewAuthHandler()
	var currentUserId int
	var loggedIn bool

	fmt.Println("Welcome to REA Game Store Cashier App")
	fmt.Println("================================")

	for{
		if !loggedIn {
            helper.ShowAuthMenu()
            choice := inputter.ReadInt("Choose option: ")

            switch choice {
            case 1:
                helper.ClearScreen()
                fmt.Println("=== User Registration ===\n")
                userID, err := authHandler.Register()
                if err != nil {
                    fmt.Printf("❌ Error: %v\n\n", err)
                } else {
                    currentUserId = userID
                    loggedIn = true
                    // Get user email for display (you might want to store this in auth handler)
                    fmt.Printf("✅ Welcome! Registration successful. User ID: %d\n\n", userID)
                }
            case 2:
                helper.ClearScreen()
                fmt.Println("=== User Login ===\n")
                userID, err := authHandler.Login()
                if err != nil {
                    fmt.Printf("❌ Error: %v\n\n", err)
                } else {
                    currentUserId = userID
                    loggedIn = true
                    fmt.Printf("✅ Login successful! Welcome back. User ID: %d\n\n", userID)
                }
            case 3:
                helper.ClearScreen()
                fmt.Println("👋 Thank you for using REA Game Store Cashier. Goodbye!")
                return
            default:
                fmt.Println("❌ Invalid choice! Please try again.\n")
            }
        }else {
            fmt.Println("You are now logged in.")
            fmt.Println("Currently assigned as User ID: %d", currentUserId)
            return
        }
	}
}
