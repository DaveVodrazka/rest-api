package server

import (
	"encoding/json"
	"net/http"

	"github.com/DaveVodrazka/rest-api/models"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	DB     *models.BooksDB
	log    *CustomLogger
}

func NewServer(r *mux.Router, db models.BooksDB) *Server {
	return &Server{router: r, DB: &db, log: &SimpleLogger}
}

func (s *Server) ListenAndServe() {
	s.log.info("Starting the server")
	http.ListenAndServe(":8000", s.router)
}

func (s *Server) Respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err = w.Write(res)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
