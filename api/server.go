package api

import (
	"net/http"
	"simple-server/storage"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store storage.Storage
	router *mux.Router
}


func (s *Server) Start() error {
	
	return http.ListenAndServe(s.listenAddr, nil)
}

func NewServer(listenAddr string, store storage.Storage) *Server{
	s := Server{
		listenAddr: listenAddr,
		store: store,
	}

	s.initRouter()
	
	return &s
}

