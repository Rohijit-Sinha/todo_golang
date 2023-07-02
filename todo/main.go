package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PostData struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	data := PostData{
		Id:    2,
		Title: "asdf",
		Body:  "Some body for you",
	}
	json.NewEncoder(w).Encode(data)
	var d PostData
	json.NewDecoder(r.Body).Decode(&d)
	// fmt.Println(d)
	url := r.URL
	query := url.Query()
	for k, v := range query {
		fmt.Printf("%v : %v, %T\n", k, v, v)
	}
}

func regidterHandlers() {
	http.HandleFunc("/", handlerfunc)
}

func main() {
	regidterHandlers()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
