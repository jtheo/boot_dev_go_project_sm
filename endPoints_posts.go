package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerGetPost(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/posts/")
	post, err := apiCfg.dbClient.GetPosts(email)
	if err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	log.Println(http.StatusOK, post)
	respondWithJSON(w, http.StatusOK, post)
}

func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	post, err := apiCfg.dbClient.CreatePost(params.UserEmail, params.Text)
	if err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)

	}
	log.Println(http.StatusCreated, post)

	respondWithJSON(w, http.StatusCreated, post)
}

func (apiCfg apiConfig) handlerUpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if err := apiCfg.dbClient.DeletePost(id); err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	log.Println(http.StatusOK, struct{}{})
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (apiCfg apiConfig) endpointPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handlerGetPost(w, r)
	case http.MethodPost:
		apiCfg.handlerCreatePost(w, r)
	case http.MethodPut:
		apiCfg.handlerUpdatePost(w, r)
	case http.MethodDelete:
		apiCfg.handlerDeletePost(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}
