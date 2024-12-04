package main

import "urlShort/internal/server"

func main() {
	s, _ := server.New()
	s.Run()
}
