all: build-server

server-app = app/server/app.go
server-dependence = internal/server/api/tcp.go
server-out = bin/server.out
build-server: $(server-app) $(server-dependence) 
	go build -o $(server-out) $(server-app)
