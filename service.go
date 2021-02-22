package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Service struct {
	Lists []ToDoList
	listIdCounter int
}

func health(w http.ResponseWriter, request *http.Request) {
	logrus.Info("API Health is OK!")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"alive": true}`)
}

func printTest(w http.ResponseWriter, r *http.Request) {
	logrus.Info("\\_Test Printed!_//")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!\n")
}

func (s *Service) createToDoList(w http.ResponseWriter, r *http.Request) {
	s.Lists = append(s.Lists, ToDoList{Id: strconv.Itoa(s.listIdCounter)})
	s.listIdCounter++
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.Lists)
}

func (s *Service) getLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.Lists)
}

func (s *Service) getList(w http.ResponseWriter, r *http.Request) {
	listId := mux.Vars(r)["id"]
	list := s.getListById(listId)
	if list.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "list with id " + listId + " not found!")
		logrus.Warn("list with id " + listId + " not found!")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func (s *Service) getItem(w http.ResponseWriter, r *http.Request) {
	listId := mux.Vars(r)["id"]
	itemId := mux.Vars(r)["itemId"]
	list := s.getListById(listId)
	if list.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "list with id " + listId + " not found!")
		logrus.Warn("list with id " + listId + " not found!")
		return
	}
	item := list.getItem(itemId)
	if item.Id == "" { // empty item
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item with id " + itemId + " not found!")
		logrus.Warn("item with id " + itemId + " not found!")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (s *Service) createItem(w http.ResponseWriter, r *http.Request) {
	listId := mux.Vars(r)["id"]
	list := s.getListById(listId)
	if list.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "list with id " + listId + " not found!")
		logrus.Warn("list with id " + listId + " not found!")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var item Item
	err := decoder.Decode(&item)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%+v\n", err)
		logrus.Error(err)
		return
	}
	list.addItem(item)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.Lists)
}

func (s *Service) getListById(id string) *ToDoList {
	for i := range s.Lists {
		if s.Lists[i].Id == id {
			return &s.Lists[i]
		}
	}
	return &ToDoList{}
}