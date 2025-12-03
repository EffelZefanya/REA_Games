package handler

import (
	"fmt"
	"strconv"
	"strings"

	"rea_games/entity"
	"rea_games/helper"
	"rea_games/repository"
)

type OrderHandler struct {
	orderRepo *repository.OrderRepository
	gameRepo  *repository.GameRepository
	inputter  *helper.Inputter
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		orderRepo: repository.NewOrderRepository(),
		gameRepo:  repository.NewGameRepository(),
		inputter:  helper.NewInputter(),
	}
}

func (h *OrderHandler) CreateOrder(userID int) error {
	fmt.Println("\n=== Available Games ===")

	games, err := h.gameRepo.GetAllGamesDisplay()
	if err != nil {
		return err
	}

	if len(games) == 0 {
		return fmt.Errorf("no games available")
	}

	for _, game := range games {
		if game.GameQuantity > 0 {
			genres, _ := h.gameRepo.GetGenresByGameID(game.GameID)
			fmt.Printf(
				"ID: %d | Title: %s | Genres: %s | Price: %.2f | Stock: %d\n",
				game.GameID,
				game.Title,
				strings.Join(genres, ", "),
				game.Price,
				game.GameQuantity,
			)
		}
	}

	gameID := h.inputter.ReadInt("\nEnter game ID: ")
	quantity := h.inputter.ReadInt("Enter quantity: ")
	if quantity <= 0 {
		fmt.Println("[Error] Quantity must be greater than 0.")
		return nil
	}

	game, err := h.gameRepo.GetGameByID(gameID)
	if err != nil {
		return fmt.Errorf("game not found")
	}

	if game.GameQuantity < quantity {
		return fmt.Errorf("not enough stock. Available: %d", game.GameQuantity)
	}

	total := game.Price * float64(quantity)

	fmt.Println("\n=== Order Summary ===")
	fmt.Println("Game    :", game.Title)
	fmt.Println("Qty     :", quantity)
	fmt.Printf("Price   : %.2f\n", game.Price)
	fmt.Printf("Total   : %.2f\n", total)

	confirm := strings.ToLower(h.inputter.ReadInput("Confirm order? (y/n): "))
	if confirm != "y" {
		fmt.Println("❌ Order cancelled.")
		return nil
	}

	order := entity.Order{
		UserID:       userID,
		GameID:       gameID,
		GameQuantity: quantity,
		TotalPrice:   total,
	}

	err = h.orderRepo.CreateOrder(&order)
	if err != nil {
		return err
	}

	newQuantity := game.GameQuantity - quantity
	err = h.gameRepo.UpdateGameQuantity(gameID, newQuantity)
	if err != nil {
		return err
	}

	fmt.Println("✅ Order created successfully!")
	return nil
}

func (h *OrderHandler) ListOrders() error {
	orders, err := h.orderRepo.GetAllOrders()
	if err != nil {
		return err
	}

	if len(orders) == 0 {
		fmt.Println("\nNo orders found.")
		return nil
	}

	fmt.Println("\n=== All Orders ===")

	for _, order := range orders {
		genres, _ := h.gameRepo.GetGenresByGameID(order.GameID)
		fmt.Printf(
			"ID: %d | User: %s | Game: %s | Genres: %s | Qty: %d | Total: %.2f | Date: %s\n",
			order.OrderID,
			order.UserEmail,
			order.GameTitle,
			strings.Join(genres, ", "),
			order.GameQuantity,
			order.TotalPrice,
			order.CreatedAt.Format("2006-01-02 15:04"),
		)
	}

	return nil
}

func (h *OrderHandler) ListUserOrders(userID int) error {
	orders, err := h.orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		return err
	}

	if len(orders) == 0 {
		fmt.Println("\nYou have no orders.")
		return nil
	}

	fmt.Println("\n=== Your Orders ===")

	for _, order := range orders {
		genres, _ := h.gameRepo.GetGenresByGameID(order.GameID)
		fmt.Printf(
			"ID: %d | Game: %s | Genres: %s | Qty: %d | Total: %.2f | Date: %s\n",
			order.OrderID,
			order.GameTitle,
			strings.Join(genres, ", "),
			order.GameQuantity,
			order.TotalPrice,
			order.CreatedAt.Format("2006-01-02"),
		)
	}

	return nil
}

func (h *OrderHandler) UpdateOrder() error {
	err := h.ListOrders()
	if err != nil {
		return err
	}

	orderID := h.inputter.ReadInt("\nEnter Order ID to update: ")

	order, err := h.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("order not found")
	}

	fmt.Printf("Current quantity: %d\n", order.GameQuantity)

	input := h.inputter.ReadInput("Enter new quantity: ")

	var newQty int

	if strings.TrimSpace(input) == "" {
		newQty = order.GameQuantity
	} else {
		parsedQty, err := strconv.Atoi(input)
		if err != nil || parsedQty <= 0 {
			return fmt.Errorf("invalid quantity")
		}
		newQty = parsedQty
	}

	game, err := h.gameRepo.GetGameByID(order.GameID)
	if err != nil {
		return fmt.Errorf("game not found")
	}

	availableStock := game.GameQuantity + order.GameQuantity

	if newQty > availableStock {
		return fmt.Errorf("not enough stock. Available: %d", availableStock)
	}

	newTotal := game.Price * float64(newQty)

	fmt.Println("\n=== Update Summary ===")
	fmt.Println("Game    :", game.Title)
	fmt.Println("Old Qty :", order.GameQuantity)
	fmt.Println("New Qty :", newQty)
	fmt.Printf("Total   : %.2f\n", newTotal)

	confirm := strings.ToLower(h.inputter.ReadInput("Confirm update? (y/n): "))
	if confirm != "y" {
		fmt.Println("❌ Update cancelled.")
		return nil
	}

	err = h.orderRepo.UpdateOrder(orderID, newQty, newTotal)
	if err != nil {
		return err
	}

	newStock := availableStock - newQty
	err = h.gameRepo.UpdateGameQuantity(order.GameID, newStock)
	if err != nil {
		return err
	}

	fmt.Println("✅ Order updated successfully!")
	return nil
}

func (h *OrderHandler) DeleteOrder() error {
	err := h.ListOrders()
	if err != nil {
		return err
	}

	idStr := h.inputter.ReadInput("\nEnter order ID to delete: ")
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid ID")
	}

	confirm := strings.ToLower(h.inputter.ReadInput("Are you sure? (y/n): "))
	if confirm != "y" {
		fmt.Println("Deletion cancelled.")
		return nil
	}

	order, err := h.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	game, err := h.gameRepo.GetGameByID(order.GameID)
	if err != nil {
		return err
	}

	newQuantity := game.GameQuantity + order.GameQuantity
	err = h.gameRepo.UpdateGameQuantity(order.GameID, newQuantity)
	if err != nil {
		return err
	}

	err = h.orderRepo.DeleteOrder(orderID)
	if err != nil {
		return err
	}

	fmt.Println("✅ Order deleted successfully!")
	return nil
}
