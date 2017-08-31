package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/littlebrownham/supermarket/adapter/marketdb"
	"github.com/littlebrownham/supermarket/endpoints"
	"github.com/urfave/negroni"
)

type Server struct {
	router     *mux.Router
	httpServer *http.Server
	middleware *negroni.Negroni

	host string
	port int
}

func (s *Server) initializeRoutes() {
	db := marketdb.NewMarketDB()
	createProduceEndpoint := endpoints.NewProduce(db)

	s.router.HandleFunc("/createProduce", createProduceEndpoint.CreateProduce).Methods("POST")
}

func New() (*Server, error) {
	s := &Server{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		host:       "localhost",
		port:       8080,
	}

	s.initializeRoutes()

	return s, nil
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.httpServer = &http.Server{Addr: addr, Handler: s.middleware}

	fmt.Printf("SuperMarket API listening on %s.....", addr)
	if err := s.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("error occurred when starting up rts %s", err)
	}

	return nil
}
