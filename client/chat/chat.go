package chat

import (
	"bufio"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func Connect(serverURL string, username string) error {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["username"] = username

	client, err := socketio_client.NewClient(serverURL, opts)
	if err != nil {
		log.Fatal("Error creating socket.io client: ", err)
	}

	for {
		Reader := bufio.NewReader(os.Stdin)
		data, err := Reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input: ", err)
		}
		client.Emit("message", data)
		log.Printf("Sended Message from %s", username)
	}

}
