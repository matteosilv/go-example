package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matteosilv/go-example/common"
)

var users = []common.User{
	{Name: "Scott", Surname: "Lang"},
	{Name: "Steve", Surname: "Rogers"},
	{Name: "James", Surname: "Barnes"}}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", userHandler)
	http.Handle("/", r)
	log.Println("listening on port 3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil || id < 0 || id > 2 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{ \"error\" : \"invalid id\"}")
		return
	}
	u := users[id]
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		log.Println(err)
	}

}
