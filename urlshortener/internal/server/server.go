package server

import (
	"urlShort/internal/api"
	"urlShort/internal/storage"
)

type Server struct {
	api *api.API
	db  *storage.DB
}

// Connetion to database
const connString = "sd"

func New() (*Server, error) {
	s := Server{}
	db, err := storage.New(connString)
	if err != nil {
		return nil, err
	}
	s.db = db
	api := api.New(s.db)
	s.api = api
	return &s, nil
}

func (s *Server) Run() {
	s.api.Run(":8080")
}
