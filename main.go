package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jtheo/socialmedia/internal/database"
)

func main() {

	const addr = "localhost:8080"
	log.Println("Starting", addr)
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", testHandler)
	serveMux.HandleFunc("/err", testErrHandler)
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if payload == nil {
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marshalling", err)
		respondWithError(w, 500, fmt.Errorf("unable to Marshal the payload"))
	}
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("don't call respondWithError with a nil err!")
		return
	}
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

type errorBody struct {
	Error string `json:"error"`
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	code := 599
	keys, ok := r.URL.Query()["err"]
	if ok {
		code, _ = strconv.Atoi(keys[0])
	}

	err := fmt.Sprintf("the error code is %v", code)
	respondWithError(w, code, fmt.Errorf(err))
}
