package server

import (
	"urlShort/internal/api"
	"urlShort/internal/config"
	"urlShort/internal/storage"
)

type Server struct {
	api    *api.API
	db     *storage.DB
	config *config.Config
}

func New() (*Server, error) {
	s := Server{}
	s.config = config.New()

	connString := s.config.DB.Username + ":" + s.config.DB.Password + "@/" + s.config.DB.Name
	db, err := storage.New(connString)
	if err != nil {
		return nil, err
	}

	s.db = db
	s.api = api.New(s.db)
	return &s, nil
}

func (s *Server) Run() {
	s.api.Run(s.config.Port)
}
