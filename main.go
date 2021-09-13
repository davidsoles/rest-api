package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Id        int32
	FirstName string
	LastName  string
}

type Users []User

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/Users", UsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/User/{name}", UserHandler).Methods(http.MethodGet)
	router.HandleFunc("/Queries", QueryHandler).Methods(http.MethodGet).Queries("page", "{page}").Queries("pageSize", "{pageSize}")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fprintf, err := fmt.Fprintf(w, "Hello David!")
	if err != nil {
		log.Println(fprintf)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	users := Users{
		User{Id: 1, FirstName: "David", LastName: "Soles"},
		User{Id: 2, FirstName: "Ruth", LastName: "Soles"},
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fprintf, err := fmt.Fprintf(w, "Your username is: %v", params["name"])
	if err != nil {
		log.Println(fprintf)
	}
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fprintf, err := fmt.Fprintf(w, "Page number is %v and pageSize is %v", params["page"], params["pageSize"])
	if err != nil {
		log.Println(fprintf)
	}
}
