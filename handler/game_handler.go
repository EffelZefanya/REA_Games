package handler

import (
	"fmt"
	"rea_games/entity"
	"rea_games/helper"
	"rea_games/repository"
	"strings"
	"time"
)

type GameHandler struct {
	gameRepo      *repository.GameRepository
	developerRepo *repository.DeveloperRepository
	genreRepo     *repository.GenreRepository
	inputter      *helper.Inputter
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		developerRepo: repository.NewDeveloperRepository(),
		gameRepo:      repository.NewGameRepository(),
		genreRepo:     repository.NewGenreRepository(),
		inputter:      helper.NewInputter(),
	}
}

func (h *GameHandler) CreateGame() error {
	fmt.Println("\n=== Add Games ===")
	var layout = "2006-01-02"
	title := h.inputter.ReadInput("Enter Title: ")
	price := h.inputter.ReadFloat("Enter Price: ")
	releasedate := h.inputter.ReadInput("Enter Date: ")
	releasedateparsed, err := time.Parse(layout, releasedate)
	if err != nil {
		fmt.Println("Error date: ", err)
		return err
	}
	description := h.inputter.ReadInput("Enter Description: ")
	developers, err := h.developerRepo.GetAllDevelopers()
	if err != nil {
		fmt.Println("database error getting developers")
		return err
	}

	for _, developer := range developers {
		fmt.Printf(
			"ID: %d | Developer Name: %s\n",
			developer.Developer_id,
			developer.Developer_Name,
		)
	}

	developer := h.inputter.ReadInt("Enter Developer: ")
	genre := []int{}

	genres, err := h.genreRepo.GetAllGenre()
	if err != nil {
		fmt.Println("database error getting genre")
		return err
	}

	for _, genre := range genres {
		fmt.Printf(
			"ID: %d | Genre Name: %s\n",
			genre.Genre_id,
			genre.Genre_Name,
		)
	}

	for {
		input := h.inputter.ReadInt("Enter Genre (enter 0 to finish): ")
		if input == 0 {
			break
		}
		genre = append(genre, input)
	}

	quantity := h.inputter.ReadInt("Enter quantity: ")
	if quantity < 0 {
		fmt.Println("Quantity must be above 0")
		return nil
	}
	Game := entity.Game{
		Title:         title,
		ReleaseDate:   releasedateparsed,
		GameQuantity:  quantity,
		Price:         price,
		DeveloperName: developer,
		Description:   description,
		Genre:         genre,
	}

	err = h.gameRepo.CreateGame(&Game)
	if err != nil {
		fmt.Println("Error Creating Game")
		return err
	}

	return nil
}

func (h *GameHandler) ListGames() error {
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

				"ID: %d\n  Title: %s\n  Genres: %s\n  Price: %.2f\n  Stock: %d\n  ReleaseDate: %s\n  Current Developer: %s\n",
				game.GameID,
				game.Title,
				genres,
				game.Price,
				game.GameQuantity,
				game.ReleaseDate.Format("2006-01-02"),
				game.DeveloperName,
			)
		}
	}
	return nil
}

func (h *GameHandler) UpdateGames() error {
	var layout = "2006-01-02"
	err := h.ListGames()
	if err != nil {
		fmt.Println("error database cannot get games")
		return err
	}
	gameID := h.inputter.ReadInt("Enter game ID to update: ")

	GameDisplay, err := h.gameRepo.GetGameDisplayByID(gameID)
	if err != nil {
		fmt.Println("Error getting game by id")
		return err
	}

	game, err := h.gameRepo.GetGameByID(gameID)
	if err != nil {
		fmt.Println("Error getting game by id")
		return err
	}

	if game.GameQuantity > 0 {
		genres, _ := h.gameRepo.GetGenresByGameID(GameDisplay.GameID)
		fmt.Printf(
			"ID: %d\n  Title: %s\n  Genres: %s\n  Price: %.2f\n  Stock: %d\n  ReleaseDate: %s\n  Current Developer: %s\n",
			GameDisplay.GameID,
			GameDisplay.Title,
			strings.Join(genres, ", "),
			GameDisplay.Price,
			GameDisplay.GameQuantity,
			GameDisplay.ReleaseDate.Format("2006-01-02"),
			GameDisplay.DeveloperName,
		)
	}

	title := h.inputter.ReadInput("Enter Title :")
	price := h.inputter.ReadFloat("Enter Price :")
	releasedate := h.inputter.ReadInput("Enter Release_Date (YYYY-MM-DD):")

	releasedateparsed, err := time.Parse(layout, releasedate)
	if err != nil {
		fmt.Println("Error date ", err)
		return err
	}
	developers, err := h.developerRepo.GetAllDevelopers()
	if err != nil {
		fmt.Println("database error getting developers")
		return err
	}

	for _, developer := range developers {
		fmt.Printf(
			"ID: %d | Developer Name: %s\n",
			developer.Developer_id,
			developer.Developer_Name,
		)
	}

	developer := h.inputter.ReadInt("Enter Developer: ")
	genre := []int{}

	genres, err := h.genreRepo.GetAllGenre()
	if err != nil {
		fmt.Println("database error getting genre")
		return err
	}

	for _, genre := range genres {
		fmt.Printf(
			"ID: %d | Genre Name: %s\n",
			genre.Genre_id,
			genre.Genre_Name,
		)
	}

	for {
		input := h.inputter.ReadInt("Enter Genre (enter 0 to finish): ")
		if input == 0 {
			break
		}
		genre = append(genre, input)
	}

	quantity := h.inputter.ReadInt("Enter quantity: ")
	if quantity < 0 {
		fmt.Println("Quantity must be above 0")
		return nil
	}

	genresearched, _ := h.gameRepo.GetGenreByID(genre)
	developersearched, _ := h.gameRepo.GetDeveloperByID(developer)
	fmt.Println("\n=== Change Summary ===")
	fmt.Printf(

		"ID: %d\n  Title: %s\n  Genres: %s\n  Price: %.2f\n  Stock: %d\n  ReleaseDate: %s\n  Current Developer: %s\n",
		game.GameID,
		title,
		genresearched,
		price,
		quantity,
		releasedateparsed.Format("2006-01-02"),
		developersearched,
	)

	confirm := strings.ToLower(h.inputter.ReadInput("Confirm order? (y/n): "))
	if confirm != "y" {
		fmt.Println("❌ Order cancelled.")
		return nil
	}

	err = h.gameRepo.UpdateGame(genre, releasedateparsed, developer, title, price, quantity, gameID)
	if err != nil {
		return err
	}

	fmt.Println("✅ game updated successfully!")
	return nil
}

func (h *GameHandler) DeleteGame() error {
	err := h.ListGames()
	if err != nil {
		fmt.Println("error database cannot get games")
		return err
	}
	gameID := h.inputter.ReadInt("Enter game ID to update: ")

	err = h.gameRepo.DeleteGame(gameID)
	if err != nil {
		return err
	}
	return nil
}
