const express = require('express');
const http = require('http');
const { Server } = require('socket.io');
const app = express();
const server = http.createServer(app);
const io = new Server(server);
const PORT = 8080;

app.get('/', (req, res) => {
    res.end('Server is running')
})
io.on('connection', (socket) => {
    socket.on('join', (username) => {
        socket.name = username;
        console.log(`Client joined as ${username} at ${new Date().toLocaleString()}`);
        io.emit('system', `[${new Date().toLocaleString()}] Client connected`);

    });

    socket.on('message', (msg) => {
        if (!socket.name) {
            socket.name = msg;
            return;
        }
        console.log(`Message received: ${msg} at ${new Date().toLocaleString()}`);
        io.emit('message', `[${new Date().toLocaleString()}] ${socket.name}: ${msg}`);
    });

    socket.on('disconnect', () => {
        console.log(`Client disconnected ${new Date().toLocaleString()}`, socket.name);
    });

});
server.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});



