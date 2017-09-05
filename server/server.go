package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/littlebrownham/supermarket/adapter/marketdb"
	"github.com/littlebrownham/supermarket/endpoints"
	"github.com/meatballhat/negroni-logrus"
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
	createProduceEndpoint := endpoints.NewCreateProduce(db)
	getProduceEndpoint := endpoints.NewGetProduce(db)
	deleteProduceEndpoint := endpoints.NewDeleteProduce(db)

	s.router.HandleFunc("/createproduce", createProduceEndpoint.CreateProduce).Methods("POST")
	s.router.HandleFunc("/getproduce", getProduceEndpoint.GetProduce).Methods("GET")
	s.router.HandleFunc("/deleteproduce", deleteProduceEndpoint.DeleteProduce).Methods("DELETE")
}

func (s *Server) initializeMiddleware() {
	s.middleware.Use(negronilogrus.NewMiddleware())

	s.middleware.UseHandler(s.router)
}

func New() *Server {
	s := &Server{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		host:       "0.0.0.0",
		port:       50200,
	}

	s.initializeRoutes()
	s.initializeMiddleware()

	return s
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.httpServer = &http.Server{Addr: addr, Handler: s.middleware}

	fmt.Printf("SuperMarket API listening on %s.....\n", addr)
	if err := s.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("error occurred when starting up rts %s", err)
	}

	return nil
}
