package Models

import (
	"time"

	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
	i "go.mongodb.org/mongo-driver/bson/primitive"
)

type OriginalModel struct {
	ID               i.ObjectID `bson:"_id,omitempty"`
	MsgModel         MsgModel
	Conn             *websocket.Conn
	ServerURL        string
	ViewModel        ViewModel
	LoginModel       LoginModel
	SendMsgModel     SendMsgModel
	ViewportUpdateCh chan string
	State            int
}

// for receiving
type MsgModel struct {
	ID        i.ObjectID `bson:"_id,omitempty"`
	UserName  string     `bson:"username"`
	Content   string     `bson:"content"`
	TimeStamp time.Time  `bson:"timestamp"`
	Type      string     `bson:"type"`
}

type Channel struct {
	ID   i.ObjectID `bson:"_id,omitempty"`
	Name string
}
type ViewModel struct {
	ID       i.ObjectID `bson:"_id,omitempty"`
	Viewport viewport.Model
}
type LoginModel struct {
	ID       i.ObjectID `bson:"_id,omitempty"`
	Username textinput.Model
	Password textinput.Model
	Block    int `default:"0"`
}
type ViewportUpdateMsg struct {
	Content string
}
type SendMsgModel struct {
	ID       i.ObjectID `bson:"_id,omitempty"`
	UserName string
	Content  textinput.Model
}

func (c Channel) Title() string       { return c.Name }
func (c Channel) Description() string { return "" }
func (c Channel) FilterValue() string { return c.Name }

func (m OriginalModel) Init() tea.Cmd {
	return m.listenForViewportUpdates()
}
func (m OriginalModel) listenForViewportUpdates() tea.Cmd {
	return func() tea.Msg {
		content := <-m.ViewportUpdateCh
		return ViewportUpdateMsg{Content: content}
	}
}

func (m OriginalModel) startMessageListener() tea.Cmd {
	return func() tea.Msg {
		go func() {
			for {
				_, msg, err := m.Conn.ReadMessage()
				if err != nil {
					fmt.Printf("\nDisconnected from server: %v\n", err)
					os.Exit(0)
				}
				var received MsgModel
				if err := json.Unmarshal(msg, &received); err != nil {
					log.Printf("Error unmarshaling message: %v\n", err)
					continue
				}
				ok, err := getAllMessages()
				if err != nil {
					log.Printf("Error fetching messages: %v\n", err)
					continue
				}
				m.ViewportUpdateCh <- ok
			}
		}()
		return nil
	}
}
