package Models

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m ViewModel) updateViewport(width, height int) ViewModel {
	m.Viewport.Width = width
	m.Viewport.Height = height
	msgs, err := getAllMessages()
	if err != nil {
		log.Fatal("Error getting messages: ", err)
	}
	for _, msg := range msgs {
		m.Content += fmt.Sprintf("[%s] %s: %s\n", msg.TimeStamp.Format("15:04"), msg.UserName, msg.Content)
	}
	m.Viewport.SetContent(m.Content)
	return m
}

func (m LoginModel) updateLogin(msg tea.Msg) LoginModel {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if strings.TrimSpace(m.Username.Value()) == "" || strings.TrimSpace(m.Password.Value()) == "" {
				fmt.Println("Username and Password cannot be empty")
				return m
			}
			//check in db
			if err := UserExists(m.Username.Value(), m.Password.Value()); err != nil {
				log.Fatal("User already exists: ", err)
			}

			if err := CreateUser(m.Username.Value(), m.Password.Value()); err != nil {
				log.Fatal("Error creating user: ", err)
			}
			fmt.Println("User created successfully!")
			return m
		}
	}
	return m
}
func (m SendMsgModel) updateSendMsg(msg tea.Msg) SendMsgModel {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if strings.TrimSpace(m.Content.Value()) == "" {
				fmt.Println("Message cannot be empty")
				return m
			}
			if err := SendMsg(m.UserName, m.Content); err != nil {
				log.Fatal("Error sending message: ", err)
			}

			m.Content.SetValue("")
		}
	}
	return m
}
