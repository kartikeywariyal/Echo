const express = require('express');
const http = require('http');
const app = express();
const server = http.createServer(app);
const WebSocket = require('ws');

const PORT = process.env.PORT ? Number(process.env.PORT) : 8081;

app.get('/', (req, res) => {
    res.end('Server is running')
})
const ws = new WebSocket.Server({ server })
ws.on('connection', (ws) => {
    ws.on('message', (raw) => {
        let msg;
        try {
            msg = JSON.parse(raw);
        } catch (e) {
            // Ignore non-JSON payloads
            return;
        }

        const userName = msg.UserName;
        const content = msg.Content;
        // First message from client is their name
        if (ws.name === undefined && typeof userName === 'string' && userName.trim() !== '') {
            ws.name = userName.trim();
            console.log(`Client joined as ${ws.name} at ${new Date().toLocaleString()}`);

            broadcast({
                Type: "join",
                UserName: ws.name,
                Content: "",
                TimeStamp: new Date().toISOString(),
            });
            return;
        }
        if (typeof content === 'string' && content.trim() !== '') {
            const effectiveName = (typeof userName === 'string' && userName.trim() !== '')
                ? userName.trim()
                : (ws.name ?? "unknown");

            console.log(`${content} at ${new Date().toLocaleString()}`);
            broadcast({
                Type: "message",
                UserName: effectiveName,
                Content: content,
                TimeStamp: new Date().toISOString(),
            });
        }
    })

    ws.on('close', () => {
        console.log(`Client ${ws.name} disconnected at ${new Date().toLocaleString()}`);
        broadcast({
            Type: "leave",
            UserName: ws.name ?? "unknown",
            Content: "",
            TimeStamp: new Date().toISOString(),
        });
    })
});
function broadcast(message) {
    const payload = (typeof message === "string") ? message : JSON.stringify(message);
    ws.clients.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.send(payload);
        }
    });
}

server.listen(PORT, 'localhost', () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});



