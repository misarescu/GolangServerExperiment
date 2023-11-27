package api

import (
	"encoding/json"
	"net/http"
	"simple-server/models"
	"strconv"

	"github.com/gorilla/mux"
)

type apiHandler func (http.ResponseWriter, *http.Request) error

func makeHandler(f apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		// handle bad requests
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}

func writeJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(data)
}

func (s *Server) initRouter() {
	s.router = mux.NewRouter()

	userRouter := s.router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{id}", makeHandler(s.handleGetUserByID)).Methods("GET")
	userRouter.HandleFunc("/{id}", makeHandler(s.handleRemoveUserById)).Methods("DELETE")
	userRouter.HandleFunc("", makeHandler(s.handleGetAllUsers)).Methods("GET")

	http.Handle("/",s.router)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		return models.NewRequestError("id type needs to be integer")
	}

	user := s.store.Get(id)
	writeJSON(w,http.StatusOK, user)
	
	return nil
}

func (s *Server) handleGetAllUsers(w http.ResponseWriter, r *http.Request) error {
	users := s.store.GetAll()
	writeJSON(w,http.StatusOK, users)

	return nil
}

func (s *Server) handleRemoveUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		return models.NewRequestError("id type needs to be integer")
	}

	user := s.store.Remove(id)
	writeJSON(w,http.StatusOK, user)

	return nil
}