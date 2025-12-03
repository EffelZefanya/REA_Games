package cli

import (
	"fmt"
	"rea_games/handler"
	"rea_games/helper"
)

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
	fmt.Println("1. 🎮 Game Management")
	fmt.Println("2. 📋 Order Management")
	fmt.Println("3. 📊 Reports")
	fmt.Println("4. ➜] Logout")
	fmt.Println("5. ⏻ Exit")
	fmt.Println("==========================")
}

func HandleOrderOperations(userID int) {
	orderHandler := handler.NewOrderHandler()
	inputter := helper.NewInputter()

	for {
		fmt.Println("\n=== Order Management ===")
		fmt.Println("1. 🛒 Create Order")
		fmt.Println("2. 📋 List All Orders")
		fmt.Println("3. 👤 List My Orders")
		fmt.Println("4. 🔄  Update Order")
		fmt.Println("5. 🗑️  Delete Order")
		fmt.Println("6. ↩️  Back to Main Menu")
		fmt.Println("========================")

		choice := inputter.ReadInt("Choose option: ")

		switch choice {
		case 1:
			fmt.Println("\n--- Create New Order ---")
			err := orderHandler.CreateOrder(userID)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 2:
			fmt.Println("\n--- All Orders ---")
			err := orderHandler.ListOrders()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 3:
			fmt.Println("\n--- My Orders ---")
			err := orderHandler.ListUserOrders(userID)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 4:
			fmt.Println("\n--- Update Order ---")
			err := orderHandler.UpdateOrder()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 5:
			fmt.Println("\n--- Delete Order ---")
			err := orderHandler.DeleteOrder()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 6:
			ClearScreen()
			return
		default:
			fmt.Println("❌ Invalid choice! Please try again.")
		}
	}
}

func HandleReportOperations() {
	reportHandler := handler.NewReportHandler()
	inputter := helper.NewInputter()

	for {
		fmt.Println("\n=== Reports ===")
		fmt.Println("1. 👥 Developers' Games report")
		fmt.Println("2. 📦 Stock Report")
		fmt.Println("3. 🛒 Revenue Report")
		fmt.Println("4. ↩️  Back to Main Menu")
		fmt.Println("================")

		choice := inputter.ReadInt("Choose option: ")

		switch choice {
		case 1:
			fmt.Println("\n--- User Report ---")
			err := reportHandler.GetDevelopersReport()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 2:
			fmt.Println("\n--- Stock Report ---")
			err := reportHandler.GetStockReport()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 3:
			fmt.Println("\n--- Revenue Report ---")
			err := reportHandler.GetRevenueReport()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 4:
			ClearScreen()
			return
		default:
			fmt.Println("❌ Invalid choice! Please try again.")
		}
	}
}

func HandleGameOperations() {
	gamehandler := handler.NewGameHandler()
	inputter := helper.NewInputter()

	for {
		fmt.Println("\n=== Game Management ===")
		fmt.Println("1. 🎮  Create Game")
		fmt.Println("2. 📋  List All Games")
		fmt.Println("3. 🔄  Update Game")
		fmt.Println("4. 🗑️  Delete Game")
		fmt.Println("5. ↩️  Back to Main Menu")
		fmt.Println("========================")

		choice := inputter.ReadInt("Choose option: ")

		switch choice {
		case 1:
			fmt.Println("\n--- Add New Game ---")
			err := gamehandler.CreateGame()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 2:
			fmt.Println("\n--- List Of Games---")
			err := gamehandler.ListGames()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 3:
			fmt.Println("\n--- Update A Game ---")
			err := gamehandler.UpdateGames()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 4:
			fmt.Println("\n--- Delete A Game ---")
			err := gamehandler.DeleteGame()
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}
		case 5:
			ClearScreen()
			return
		default:
			fmt.Println("❌ Invalid choice! Please try again.")
		}
	}
}
