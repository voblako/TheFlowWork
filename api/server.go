package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/voblako/TheFlowWork/storage"
)

type Config struct {
	ListenAddr string
	Storage storage.Config
}

type Server struct {
	*Config
	Started time.Time
	DB storage.DB
}

func NewServer(config Config) (*Server, error) {
	db, err := storage.NewPostgressConnect(config.Storage)
	if err!=nil{
		return &Server{}, err
	}
	return &Server{
		Config: &config,
		DB: db,
	}, nil
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.Health)

	server := http.Server{
		Addr:    s.Config.ListenAddr,
		Handler: mux,
	}

	s.Started = time.Now().UTC()
	slog.Info("api is running", "address", s.Config.ListenAddr)
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("API stopped", "error", err)
	}
}
