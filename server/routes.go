package server

func (s *Server) Routes() {
	s.router.HandleFunc("/api/books", s.GetBooks).Methods("GET")
	s.router.HandleFunc("/api/books/{id}", s.GetBook).Methods("GET")
	s.router.HandleFunc("/api/books", s.CreateBook).Methods("POST")
	s.router.HandleFunc("/api/books/{id}", s.UpdateBook).Methods("PUT")
	s.router.HandleFunc("/api/books/{id}", s.DeleteBook).Methods("DELETE")
}
