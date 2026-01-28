package Models

import (
	db "Echo/client/db"
	"context"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/gorilla/websocket"
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
func CreateUser(username string, password string) error {
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

func SendMsg(username string, content textinput.Model) error {
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
func SendRecieveMsg(msg MsgModel) error {
	collection := db.GetCollection("Echo", "Messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	msgData := map[string]interface{}{
		"username":  msg.UserName,
		"content":   msg.Content,
		"timestamp": msg.TimeStamp,
		"type":      msg.Type,
	}
	_, err := collection.InsertOne(ctx, msgData)
	return err
}

func getAllMessages() (string, error) {
	collection := db.GetCollection("Echo", "Messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return "", err
	}
	defer data.Close(ctx)

	var messages []MsgModel
	for data.Next(ctx) {
		var msg MsgModel
		if err := data.Decode(&msg); err != nil {
			continue
		}
		messages = append(messages, msg)

	}
	msgs := ""
	for _, m := range messages {
		msgs += `[` + m.TimeStamp.Format("15:04:05") + `] ` + m.UserName + `: ` + m.Content + `
`
	}
	return msgs, nil
}

func InitialModel(serverURL string) OriginalModel {
	LoginModel := new(LoginModel)
	UsernameInput := textinput.New()
	UsernameInput.Placeholder = "Enter your username"
	UsernameInput.Focus()
	LoginModel.Username = UsernameInput
	PasswordInput := textinput.New()
	PasswordInput.Placeholder = "Enter your password"
	LoginModel.Password = PasswordInput

	ViewModel := new(ViewModel)
	Viewport := viewport.New(80, 20)
	ViewModel.Viewport = Viewport

	SendMsgModel := new(SendMsgModel)
	ContentArea := textinput.New()
	ContentArea.Placeholder = "Type your message here..."
	ContentArea.Focus()
	SendMsgModel.Content = ContentArea

	return OriginalModel{
		ServerURL:        serverURL,
		LoginModel:       *LoginModel,
		ViewModel:        *ViewModel,
		SendMsgModel:     *SendMsgModel,
		State:            1,
		ViewportUpdateCh: make(chan string, 10),
	}
}
func Wbconnect(serverURL string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
