package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"Echo/client/chat"
)

func main() {
	// Ask user for username
	Reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")

	username, err := Reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading username: ", err)
	}
	username = strings.TrimSpace(username)

	if username == "" {
		log.Fatal("Username cannot be empty")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	serverURL := "ws://localhost:" + port

	fmt.Printf("Connecting to server at %s...\n", serverURL)

	// Connect to the server and start chat
	if err := chat.Connect(serverURL, username); err != nil {
		log.Fatal("Failed to connect: ", err)
	}

}
