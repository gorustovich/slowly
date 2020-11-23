package app

import (
	"net/http"
	"log"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	server := &Server{
		mux: http.NewServeMux(),
	}
	return server
}

func (a *Server) Setup() {
	slowProcessor := NewSlowProcessor()

	a.mux.Handle("/api/slow",
		ValidateJsonType(
			CancelTooLong(
				NewSlowHandler(slowProcessor),
			),
		),
	)
}

func (a *Server) Start() {
	log.Println("Start on on :8080...")
	err := http.ListenAndServe(":8080", a.mux)
	log.Fatal(err)
}
