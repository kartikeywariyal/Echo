# Echo 💬

#### A sleek, modern terminal chat application with real-time messaging

Echo is a terminal-user-interface (TUI) chat application for lightweight and fast chatting, designed to run directly from the terminal with zero distractions. Built for developers who live in the terminal and want beautiful, real-time chat without leaving their workflow.

**Features:**
- ✨ **Real-time WebSocket messaging** - Instant message delivery
- 🎨 **Beautiful TUI UI** - Styled with lipgloss for a modern look
- 💬 **Chat bubbles** - Left/right aligned messages like modern chat apps
- 🔐 **User authentication** - Login system with MongoDB
- 📨 **Message persistence** - All messages stored in MongoDB
- ⚡ **Fast & lightweight** - No bloat, pure terminal experience

### Tech Stack 

#### Server 
- Node.js + Express.js
- WebSocket (ws library)
- Mongoose + MongoDB Atlas 

#### Client 
- Go 1.21+
- Bubble Tea (TUI framework)
- Lipgloss (styling)
- Gorilla WebSocket
    
#### Database 
- MongoDB Atlas 
  - **Users collection** - Username & password
  - **Messages collection** - Sender, content, timestamp

### Project Structure

```
Echo/
├── README.md                 <- You are here
├── CONTRIBUTING.md           <- Contribution guidelines
├── package.json              <- Server dependencies
├── go.mod                    <- Go dependencies
│ 
├── client/                   <- Go TUI Client
│   ├── main.go              <- Entry point
│   ├── chat/                <- WebSocket connection
│   │   └── chat.go
│   ├── db/                  <- Database operations
│   │   └── db.go
│   └── models/              <- UI & data models
│       ├── model.go         <- Core model structure
│       ├── initialModel.go  <- Model initialization
│       ├── update.go        <- Update logic & event handling
│       └── views.go         <- UI rendering with lipgloss
│
└── server/                  <- Node.js Chat Server
    └── server.js            <- WebSocket server & API
```

### Features in Detail

#### 🎨 Modern UI
- Rounded borders with gradient colors (cyan, magenta, green)
- Emoji indicators (✨, 💬, 📨, ✍️, 👤, 🔐)
- Dark theme with good contrast
- Responsive to terminal size

#### 💬 Smart Chat Display
- **Your messages** → Blue bubble, right-aligned 🔵
- **Others' messages** → Pink bubble, left-aligned 💗
- Timestamp for each message `[HH:MM]`
- Username display
- Scrollable message history

#### ⚡ Real-time Updates
- No polling - true event-driven updates
- WebSocket connection stays open
- Goroutine-based message listener
- Channel-based async UI updates using Bubble Tea commands

### Quick Start 

#### Prerequisites
- **Node.js** (v14+) - for the server
- **Go** (v1.21+) - for the client
- **MongoDB Atlas** - create a free account and get connection string
- **Environment variables** - create `.env` file

#### Environment Setup

Create `.env` file in project root:
```env
DB_CONNECTION_STRING=mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority
PORT=8080
```

#### Server Setup

```bash
cd server
npm install express ws cors
npm install --save-dev nodemon
npm start
```

Server runs on `ws://localhost:8080`

#### Client Setup

```bash
cd client
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver
go get github.com/joho/godotenv

go run .
```

### How to Use

1. **Launch** → `go run .` in `/client`
2. **Login** → Enter username and password (auto-creates account if doesn't exist)
3. **Chat** → Type messages and press ENTER
4. **Navigate** → Press TAB to switch between message view and input
5. **Exit** → Press CTRL+C to quit

### Architecture

#### Message Flow
```
Client Types Message
    ↓
Sent via WebSocket to Server
    ↓
Server stores in MongoDB
    ↓
Server broadcasts to all connected clients
    ↓
Client receives via goroutine listener
    ↓
Sent through channel to Update()
    ↓
UI re-renders with new message
```

#### Real-time Updates (Non-blocking)
- **Goroutine**: Listens on WebSocket continuously
- **Channel**: Passes messages to UI thread safely
- **Tea Command**: Processes channel data and triggers Update()
- **Batch**: Runs listener + keyboard input in parallel

### Contact & Support

For issues, suggestions, or contributions:
- Check [CONTRIBUTING.md](./CONTRIBUTING.md)
- Discord: Washikiballa-San

### License

MIT License - See LICENSE file for details

---

**Happy Chatting! 🎉**

