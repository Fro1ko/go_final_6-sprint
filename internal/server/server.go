package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	srv := &Server{
		Logger: logger,
		Server: &http.Server{
			Addr:         ":8080",           
			Handler:      mux,               
			ErrorLog:     logger,           
			ReadTimeout:  5 * time.Second,   
			WriteTimeout: 10 * time.Second,  
			IdleTimeout:  15 * time.Second,  
	},
}
	mux.HandleFunc("/", srv.IndexHandler)
	mux.HandleFunc("/upload", srv.UploadHandler)

	return srv
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	handlers.IndexHandler(w, r)
}

func (s *Server) UploadHandler(w http.ResponseWriter, r *http.Request) {
	handlers.UploadHandler(w, r)
}