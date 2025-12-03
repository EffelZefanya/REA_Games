package handler

import (
	"fmt"
	"rea_games/helper"
	"rea_games/repository"
)

type ReportHandler struct {
	reportRepo *repository.ReportRepository
	inputter   *helper.Inputter
}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{
		reportRepo: repository.NewReportRepository(),
		inputter:   helper.NewInputter(),
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
