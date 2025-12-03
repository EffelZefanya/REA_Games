package handler

import (
	"fmt"
	"rea_games/helper"
	"rea_games/repository"
)

type ReportHandler struct {
	reportRepo *repository.ReportRepository
	orderRepo *repository.OrderRepository
	gameRepo *repository.GameRepository
	inputter  *helper.Inputter
}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{
		reportRepo: repository.NewReportRepository(),
		gameRepo: repository.NewGameRepository(),
		orderRepo: repository.NewOrderRepository(),
		inputter:  helper.NewInputter(),
	}
}

func (h *ReportHandler) GetDevelopersReport() error {
	// Call the repository function to fetch the report data
	reportData, err := h.reportRepo.GetDevelopersReport()
	if err != nil {
		return err
	}

	// Check if any data was returned
	if len(reportData) == 0 {
		return err
	}

	// Print the report data
	fmt.Println("--- 📈 Developers Report ---")
	for _, report := range reportData {
		fmt.Printf(
			"👤 Developer: %s\n"+
				"    🎮 Total Games: %d\n"+
				"    📝 List of Games: %s\n",
			report.DeveloperName,
			report.TotalGamesDeveloped,
			report.ListOfGames,
		)
		fmt.Println("---------------------------")
	}
	return nil
}

func (h *ReportHandler) GetStockReport() error {
    games, err := h.gameRepo.GetAllGamesDisplay()
    if err != nil {
        return err
    }

    fmt.Println("\n=== Stock Report ===")
    var totalValue float64
    var lowStock []string

    for _, game := range games {
        value := game.Price * float64(game.GameQuantity)
        totalValue += value
        
        fmt.Printf("Title: %s | Developer: %s | Quantity: %d | Price: $%.2f | Value: $%.2f\n", 
            game.Title, game.DeveloperName, game.GameQuantity, game.Price, value)
        
        if game.GameQuantity < 5 {
            lowStock = append(lowStock, game.Title)
        }
    }

    fmt.Printf("\nTotal Inventory Value: $%.2f\n", totalValue)
    if len(lowStock) > 0 {
        fmt.Println("Low Stock Items:")
        for _, item := range lowStock {
            fmt.Printf("  - %s\n", item)
        }
    }
    return nil
}

func (h *ReportHandler) GetRevenueReport() error {
    orders, err := h.orderRepo.GetAllOrders()
    if err != nil {
        return err
    }

    fmt.Println("\n=== Order Report ===")
    var totalRevenue float64
    orderCount := len(orders)

    for _, order := range orders {
        totalRevenue += order.TotalPrice
    }

    fmt.Printf("Total Orders: %d\n", orderCount)
    fmt.Printf("Total Revenue: $%.2f\n", totalRevenue)
    fmt.Printf("Average Order Value: $%.2f\n", totalRevenue/float64(orderCount))
    
    return nil
}