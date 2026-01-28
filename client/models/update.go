package Models

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m OriginalModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	var vpCmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ViewModel.Viewport.Width = msg.Width
		if msg.Height > 12 {
			m.ViewModel.Viewport.Height = msg.Height - 12
		}
		m.ViewModel.Viewport, vpCmd = m.ViewModel.Viewport.Update(msg)
		return m, tea.Batch(vpCmd, m.listenForViewportUpdates())
	case ViewportUpdateMsg:
		m.ViewModel.Viewport.SetContent(msg.Content)
		m.ViewModel.Viewport.GotoBottom()
		return m, m.listenForViewportUpdates()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			if m.Conn != nil {
				m.Conn.Close()
			}
			return m, tea.Quit
		case "tab":
			if m.State == 1 {
				m.LoginModel.Block = (m.LoginModel.Block + 1) % 2
				if m.LoginModel.Block == 0 {
					m.LoginModel.Username.Focus()
					m.LoginModel.Password.Blur()
				} else {
					m.LoginModel.Username.Blur()
					m.LoginModel.Password.Focus()
				}
			}
			if m.State == 2 {
				m.State = 3
				m.SendMsgModel.Content.Focus()
			}
			if m.State == 3 {
				m.State = 2
				m.SendMsgModel.Content.Blur()
			}
		case "enter":
			if m.State == 1 {
				username := m.LoginModel.Username.Value()
				password := m.LoginModel.Password.Value()

				if err := UserExists(username, password); err == nil {
					if passwordCorrect, _ := ValidateUser(username, password); !passwordCorrect {
						log.Fatal("Incorrect password. Please try again.")
						return m, nil
					}
				} else {
					err := CreateUser(username, password)
					if err != nil {
						log.Fatal("Error creating user:", err)
						return m, nil
					}
				}

				conn, err := Wbconnect(m.ServerURL)
				m.Conn = conn
				m.SendMsgModel.UserName = username
				if err != nil {
					log.Fatal("Error creating user:", err)
				}

				m.State = 2
				msg, err := getAllMessages()
				if err != nil {
					log.Fatal("Error fetching messages:", err)
				}
				m.ViewModel.Viewport.SetContent(msg)
				m.ViewModel.Viewport.GotoBottom()
				return m, tea.Batch(m.listenForViewportUpdates(), m.startMessageListener())
			}
			if m.State == 2 {
				msgContent := m.SendMsgModel.Content.Value()
				if msgContent != "" {
					err := SendMsg(m.SendMsgModel.UserName, m.SendMsgModel.Content)
					if err != nil {
						log.Println("Error sending message:", err)
					}
					m.SendMsgModel.Content.SetValue("")
				}
				msg := MsgModel{
					UserName:  m.SendMsgModel.UserName,
					Content:   msgContent,
					TimeStamp: time.Now(),
					Type:      "message",
				}
				if err := m.Conn.WriteJSON(msg); err != nil {
					log.Fatal("Error sending message via websocket:", err)
				}
				msgGet, _ := getAllMessages()
				m.ViewModel.Viewport.SetContent(msgGet)
				m.ViewModel.Viewport.GotoBottom()
			}
		}
	}

	if m.State == 1 {
		m.LoginModel.Username, cmd = m.LoginModel.Username.Update(msg)
		m.LoginModel.Password, cmd = m.LoginModel.Password.Update(msg)
	}
	if m.State == 2 {
		m.SendMsgModel.Content, cmd = m.SendMsgModel.Content.Update(msg)
		m.ViewModel.Viewport, vpCmd = m.ViewModel.Viewport.Update(msg)
	}
	return m, tea.Batch(cmd, vpCmd, m.listenForViewportUpdates())
}
