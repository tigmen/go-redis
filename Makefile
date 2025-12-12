all: build-server build-client

server-app = app/server/app.go
server-dependence = internal/server/api/tcp.go
server-out = bin/server.out
build-server: $(server-app) $(server-dependence) 
	go build -o $(server-out) $(server-app)

client-app = app/client/app.go
client-dependence =
client-out = bin/client.out
build-client: $(client-app) $(client-dependence)
	go build -o $(client-out) $(client-app)
