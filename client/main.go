package main

import (
	db "Echo/client/db"
	Models "Echo/client/models"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	db.ConnectMongo(os.Getenv("DB_CONNECTION_STRING"))
	serverURL := "ws://localhost:" + port
	Model := Models.InitialModel(serverURL)
	p := tea.NewProgram(Model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}
