package server

import (
	"servcached/pkg/api"
	"servcached/pkg/config"
	"servcached/pkg/storage"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	api    *api.API
	db     *redis.Client
	config *config.APIConfig
}

func New() (*Server, error) {
	s := Server{}
	s.config = config.New()

	s.db, _ = storage.New(s.config)

	s.api = api.New(s.db)
	return &s, nil
}

func (s *Server) Run() {
	s.api.Run(s.config.Port)
}
