# Socket Chat
This is a small chat app using web sockets to learn Go.

## Getting started
First of all, you need to start a server. You can do it by running
`go run cmd/sc/main.go server`

It will start a server on `localhost`, with a default port of `3000` (run `-h` to change those settings).
Once the server is running, attach a client by running
`go run cmd/sc/main.go client`

Happy chatting!
