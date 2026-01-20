package Models

import (
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	i "go.mongodb.org/mongo-driver/bson/primitive"
)

type OriginalModel struct {
	ID           i.ObjectID `bson:"_id,omitempty"`
	MsgModel     MsgModel
	ViewModel    ViewModel
	LoginModel   LoginModel
	SendMsgModel SendMsgModel
}
type MsgModel struct {
	ID        i.ObjectID `bson:"_id,omitempty"`
	UserName  string
	Content   string
	TimeStamp time.Time
	Type      string
}

type Channel struct {
	ID   i.ObjectID `bson:"_id,omitempty"`
	Name string
}
type ViewModel struct {
	ID       i.ObjectID `bson:"_id,omitempty"`
	Content  string
	Ready    bool `default:"false"`
	Viewport viewport.Model
}
type LoginModel struct {
	ID       i.ObjectID `bson:"_id,omitempty"`
	Username textarea.Model
	Password textarea.Model
}
type SendMsgModel struct {
	ID       i.ObjectID `bson:"_id,omitempty"`
	UserName string
	Content  textarea.Model
}

func (c Channel) Title() string       { return c.Name }
func (c Channel) Description() string { return "" }
func (c Channel) FilterValue() string { return c.Name }

func (m OriginalModel) Init() tea.Cmd {
	return nil
}
