package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jtheo/socialmedia/internal/database"
)

type errorBody struct {
	Error string `json:"error"`
}

type apiConfig struct {
	dbClient database.Client
}

func New(addr, db string) {
	c := database.NewClient(db)
	c.EnsureDB()
	apiCfg := apiConfig{dbClient: c}

	serveMux := http.NewServeMux()

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

	log.Println("Starting", addr)
	log.Fatal(srv.ListenAndServe())
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
		respondWithError(w, http.StatusInternalServerError, fmt.Errorf("unable to Marshal the payload"))
		return
	}
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("don't call respondWithError with a nil err!")
		return
	}
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}
