package main

import (
	"github.com/DaveVodrazka/rest-api/persistence"
	"github.com/DaveVodrazka/rest-api/server"
	"github.com/gorilla/mux"
)

func main() {
	s := server.NewServer(mux.NewRouter(), persistence.DB)
	s.Routes()
	s.ListenAndServe()
}
