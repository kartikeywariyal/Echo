package chat

import (
	Models "Echo/client/models"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func Connect(serverURL string, username string) error {

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		return fmt.Errorf("websocket dial error: %v", err)
	}
	//fmt.Println("Connected to server")

	defer conn.Close()

	joinMsg := Models.MsgModel{
		UserName:  username,
		Content:   fmt.Sprintf("%s has joined the chat", username),
		TimeStamp: time.Now(),
		Type:      "message",
	}
	if err := conn.WriteJSON(joinMsg); err != nil {
		return fmt.Errorf("failed to send join message: %v", err)
	}
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("\nDisconnected from server: %v\n", err)
				os.Exit(0)
			}
			var received Models.MsgModel
			if err := json.Unmarshal(msg, &received); err != nil {
				continue
			}
			timestamp := received.TimeStamp.Format("15:04:05")
			switch received.Type {
			case "join":
				fmt.Printf("\n[%s] %s joined the chat\n", timestamp, received.UserName)
			case "leave":
				fmt.Printf("\n[%s] %s left the chat\n", timestamp, received.UserName)
			case "message":
				fmt.Printf("\n[%s] %s: %s\n", timestamp, received.UserName, received.Content)
			default:
				fmt.Printf("\n[%s] %s: %s\n", timestamp, received.UserName, received.Content)
			}
			fmt.Print("> ")
		}
	}()

	fmt.Println("Type your messages below (Ctrl+C to exit):")
	scanner := bufio.NewScanner(os.Stdin)
	for {

		if !scanner.Scan() {
			break
		}

		content := strings.TrimSpace(scanner.Text())
		if content == "" {
			continue
		}

		msg := Models.MsgModel{
			UserName:  username,
			Content:   content,
			TimeStamp: time.Now(),
			Type:      "message",
		}

		if err := conn.WriteJSON(msg); err != nil {
			return fmt.Errorf("failed to send message: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	return nil
}
