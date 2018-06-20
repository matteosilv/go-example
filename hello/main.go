package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matteosilv/go-example/common"
)

var userService = common.GetEnv("USER_SERVICE", "http://localhost:3000")

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{id}", helloHandler)
	http.Handle("/", r)
	log.Println("listening on port 3000")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	resp, err := http.Get(userService + "/user/" + id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Can't contact the user service")
		return
	}
	if resp.StatusCode != 200 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid id %v\n", id)
		return
	}
	var u common.User
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Unespected error")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %v\n", u.Name)
}
