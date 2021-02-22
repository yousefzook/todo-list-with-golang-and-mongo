package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RestController struct {
	r *mux.Router
	s Service
}

func (c *RestController) init() {
	c.r = mux.NewRouter()
	s := Service{}
	c.r.HandleFunc("/", printTest).Methods("GET")
	c.r.HandleFunc("/health", health).Methods("GET")
	c.r.HandleFunc("/lists", s.getLists).Methods("GET")
	c.r.HandleFunc("/lists/{id}", s.getList).Methods("GET")
	c.r.HandleFunc("/lists/create", s.createToDoList).Methods("POST")
	c.r.HandleFunc("/lists/{id}/create-item", s.createItem).Methods("POST")
	c.r.HandleFunc("/lists/{id}/{itemId}", s.getItem).Methods("GET")
}

func (c *RestController) run(port string) {
	logrus.Info("Controller Running on port " + port + "...")
	handler := cors.New(cors.Options{AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"}}).Handler(c.r)
	http.ListenAndServe(port, handler)
}
