package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jtheo/socialmedia/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

func main() {
	const addr = "localhost:8080"

	c := database.NewClient("./db.json")
	c.EnsureDB()
	apiCfg := apiConfig{dbClient: c}

	log.Println("Starting", addr)
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", testHandler)

	serveMux.HandleFunc("/users", apiCfg.endpointUsersHandler)
	serveMux.HandleFunc("/users/", apiCfg.endpointUsersHandler)

	serveMux.HandleFunc("/posts", apiCfg.endpointPostHandler)
	serveMux.HandleFunc("/posts/", apiCfg.endpointPostHandler)
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
	// log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

type errorBody struct {
	Error string `json:"error"`
}
