package main

import (
	"log"
	"urlShort/internal/server"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env files")
	}
}

func main() {
	s, _ := server.New()
	s.Run()
}
