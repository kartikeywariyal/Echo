package Models

import (
	db "Echo/client/db"
	"context"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
)

func UserExists(username, password string) error {
	collection := db.GetCollection("Echo", "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := map[string]interface{}{
		"username": username,
	}
	var result map[string]interface{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}
func CreateUser(username, password string) error {
	collection := db.GetCollection("Echo", "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := map[string]interface{}{
		"username": username,
		"password": password,
	}
	_, err := collection.InsertOne(ctx, user)
	return err
}

func SendMsg(username string, content textarea.Model) error {
	collection := db.GetCollection("Echo", "Messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	msg := map[string]interface{}{
		"username":  username,
		"content":   content.Value(),
		"timestamp": time.Now(),
		"type":      "text",
	}
	_, err := collection.InsertOne(ctx, msg)
	return err
}
func getAllMessages() ([]MsgModel, error) {
	collection := db.GetCollection("Echo", "Messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer data.Close(ctx)

	var messages []MsgModel
	for data.Next(ctx) {
		var msg MsgModel
		if err := data.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)

	}
	return messages, nil
}
