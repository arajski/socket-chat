# socket-chat
This is a small chat app using web sockets to learn Go.

## Getting started
### Installation
First build the project by running
```bash
go build ./cmd/sc
```

### Running a server
To start handling chat clients, we need a running server. Start it up by running
```bash
sc server [-hostname] [-port]
```
By default, server is running on `localhost`, port `3000`

### Running a client
Once a server is running, a separate clients can be attached by running:

```bash
sc client [-hostname] [-port]
```
Default settings are also `localhost` and port `3000`

Happy chatting!

## License

[MIT](https://github.com/arajski/socket-chat/blob/main/LICENSE.md)

