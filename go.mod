module github.com/arajski/socket-chat

go 1.21

require internal/server v1.0.0
replace internal/server => ./internal/server
require internal/client v1.0.0
replace internal/client => ./internal/client
