package api

import (
	"encoding/json"
	"errors"
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
			writeJSON(w, http.StatusBadRequest, err)
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
	userRouter.HandleFunc("/{id}", makeHandler(s.handleUpdateUserById)).Methods("PUT")
	userRouter.HandleFunc("", makeHandler(s.handleGetAllUsers)).Methods("GET")
	userRouter.HandleFunc("", makeHandler(s.handleUpdateMultipleUsers)).Methods("PUT")

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

func (s *Server) handleUpdateUserById(w http.ResponseWriter, r *http.Request) error {
	if headerContentType := r.Header.Get("Content-Type"); headerContentType != "application/json" {
		return models.NewRequestError("Content type is not json!")
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		return models.NewRequestError("id type needs to be integer")
	}
	var user models.User

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&user); err != nil{
		var unmarshallErr *json.UnmarshalTypeError

		if errors.As(err, &unmarshallErr) {
			return models.NewRequestError("Bad Request: Wrong type provided")
		} else {
			return models.NewRequestError("Bad Request: " + err.Error())
		}
	}

	retUser := s.store.Update(id, &user)

	writeJSON(w,http.StatusOK, retUser)

	return nil
}

func (s *Server) handleUpdateMultipleUsers(w http.ResponseWriter, r *http.Request) error {
	if headerContentType := r.Header.Get("Content-Type"); headerContentType != "application/json" {
		return models.NewRequestError("Content type is not json!")
	}

	var users []*models.User

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&users); err != nil{
		var unmarshallErr *json.UnmarshalTypeError

		if errors.As(err, &unmarshallErr) {
			return models.NewRequestError("Bad Request: Wrong type provided")
		} else {
			return models.NewRequestError("Bad Request: " + err.Error())
		}
	}

	var result []*models.User

	for _, usr := range users {
		result = append(result, s.store.Update(usr.Id, usr))
	}


	writeJSON(w,http.StatusOK, result)

	return nil
}