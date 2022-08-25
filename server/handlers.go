package server

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/DaveVodrazka/rest-api/models"
	"github.com/gorilla/mux"
)

func (s *Server) GetBooks(w http.ResponseWriter, r *http.Request) {
	s.Respond(w, r, s.DB, http.StatusOK)
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range *s.DB {
		if item.ID == params["id"] {
			s.Respond(w, r, item, http.StatusOK)
			return
		}
	}

	s.Respond(w, r, models.Book{}, http.StatusNotFound)
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := s.Decode(r, &book)

	if err != nil {
		s.Respond(w, r, models.Book{}, http.StatusBadRequest)
	}

	book.ID = strconv.Itoa(rand.Intn(1000000))
	*s.DB = append(*s.DB, book)
	s.Respond(w, r, book, http.StatusOK)
}

func (s *Server) UpdateBook(w http.ResponseWriter, r *http.Request) {
	s.Respond(w, r, models.Book{}, http.StatusNoContent)
}

func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {
	s.Respond(w, r, models.Book{}, http.StatusNoContent)
}
