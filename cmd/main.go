package main

import (
	"Shortify/internal/mongodb"
	"Shortify/internal/server"
)

func main() {
	mongodb.InitMongo()
	server.StartServer()
}
