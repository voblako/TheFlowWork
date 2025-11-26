package api

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/voblako/TheFlowWork/internal/config"
	"github.com/voblako/TheFlowWork/storage"
)

type Server struct {
	Config  config.Config
	Started time.Time
	DB      storage.DB
}

func NewServer(config config.Config) (Server, error) {

	DatabaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Domain,
		config.Database.Port,
		config.Database.DBName)

	db, err := storage.NewPostgressConnect(storage.Config{DatabaseURL: DatabaseUrl})

	if err != nil {
		return Server{}, err
	}
	return Server{
		Config: config,
		DB:     db,
	}, nil
}

func (s *Server) Start() {
	_, err := s.DB.Init()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.Health)
	mux.HandleFunc("POST /user", s.CreateUser)
	mux.HandleFunc("GET /user/{user_id}", s.GetUser)

	wrappedMux := NewLogger(mux)

	s.Started = time.Now().UTC()
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: wrappedMux,
	}

	slog.Info("api is running", "address", fmt.Sprintf(":%s", s.Config.Port))
	err = server.ListenAndServe()
	if err != nil {
		slog.Error("API stopped", "error", err)
	}
}
