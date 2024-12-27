package main

import (
	"log"
	"servcached/pkg/server"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env files")
	}
}

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal("ssss")
	}
	s.Run()
}
