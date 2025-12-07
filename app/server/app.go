package main

import api "main/internal/server/api" 

func main() {
	server := api.NewServer(":8080")
	server.Start()
}
