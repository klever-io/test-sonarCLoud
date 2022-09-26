package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/klever-io/gcp-logging/server"
)

func main() {
	server.Serve()
}
