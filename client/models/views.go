package Models

import (
	db "Echo/client/db"
	"context"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	primaryColor  = lipgloss.Color("39")  // Cyan
	accentColor   = lipgloss.Color("135") // Magenta
	successColor  = lipgloss.Color("42")  // Green
	errorColor    = lipgloss.Color("196") // Red
	bgColor       = lipgloss.Color("234") // Dark gray
	textColor     = lipgloss.Color("255") // White
	mutedColor    = lipgloss.Color("240") // Gray
	ownMsgColor   = lipgloss.Color("38")  // Blue
	otherMsgColor = lipgloss.Color("176") // Pink

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Background(bgColor).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(accentColor).
			Align(lipgloss.Center)

	inputContainerStyle = lipgloss.NewStyle().
				Padding(1).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(primaryColor).
				Background(bgColor).
				Foreground(textColor)

	messagesContainerStyle = lipgloss.NewStyle().
				Padding(1).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(successColor).
				Background(lipgloss.Color("235"))

	ownMessageStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Background(ownMsgColor).
			Padding(0, 1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(ownMsgColor).
			Align(lipgloss.Right).
			MaxWidth(50)

	otherMessageStyle = lipgloss.NewStyle().
				Foreground(textColor).
				Background(otherMsgColor).
				Padding(0, 1).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(otherMsgColor).
				Align(lipgloss.Left).
				MaxWidth(50)

	labelStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	footerStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			Align(lipgloss.Center).
			Padding(1, 0)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Background(bgColor).
			Padding(0, 1).
			Bold(true)

	usernameStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true)
)

func (m OriginalModel) formatMessageContent() string {
	collection := db.GetCollection("Echo", "Messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return ""
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

	var formatted []string
	for _, msg := range messages {
		timeStr := msg.TimeStamp.Format("15:04")
		isOwn := msg.UserName == m.SendMsgModel.UserName

		contentLine := "[" + timeStr + "] " + msg.UserName + ":\n" + msg.Content

		if isOwn {

			styled := ownMessageStyle.Render(contentLine)
			formatted = append(formatted, lipgloss.Place(80, lipgloss.Height(styled), lipgloss.Right, lipgloss.Center, styled))
		} else {

			styled := otherMessageStyle.Render(contentLine)
			formatted = append(formatted, lipgloss.Place(80, lipgloss.Height(styled), lipgloss.Left, lipgloss.Center, styled))
		}
		formatted = append(formatted, "")
	}

	return strings.Join(formatted, "\n")
}

func (m OriginalModel) LoginView() string {
	title := titleStyle.Render("✨ Echo Chat ✨")

	usernameLabel := labelStyle.Render("👤 Username:")
	usernameInput := inputContainerStyle.Render(m.LoginModel.Username.View())

	passwordLabel := labelStyle.Render("🔐 Password:")
	passwordInput := inputContainerStyle.Render(m.LoginModel.Password.View())

	hints := footerStyle.Render("💡 Use TAB to switch • ENTER to login • CTRL+C to quit")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"\n",
		usernameLabel,
		usernameInput,
		"\n",
		passwordLabel,
		passwordInput,
		"\n",
		hints,
	)

	return content
}

func (m OriginalModel) mainView() string {

	header := statusBarStyle.Render("💬 Echo Chat Room")

	messagesLabel := labelStyle.Render("📨 Messages")
	formattedMsgs := m.formatMessageContent()
	messagesContent := messagesContainerStyle.Width(80).Render(formattedMsgs)
	messagesSection := lipgloss.JoinVertical(
		lipgloss.Top,
		messagesLabel,
		messagesContent,
	)

	inputLabel := labelStyle.Render("✍️  Type your message:")
	inputContent := inputContainerStyle.Render(m.SendMsgModel.Content.View())
	inputSection := lipgloss.JoinVertical(
		lipgloss.Top,
		inputLabel,
		inputContent,
	)

	hints := footerStyle.Render("💡 TAB to focus message • ENTER to send • CTRL+C to quit")

	mainContent := lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		"\n",
		messagesSection,
		"\n",
		inputSection,
		"\n",
		hints,
	)

	return mainContent
}

func (m OriginalModel) View() string {
	if m.State == 1 {
		return m.LoginView()
	}
	if m.State == 2 {
		return m.mainView()
	}
	return "Unknown State"
}
