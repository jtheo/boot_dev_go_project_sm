package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handlerGetUser(w, r)
	case http.MethodPost:
		apiCfg.handlerCreateUser(w, r)
	case http.MethodPut:
		apiCfg.handlerUpdateUser(w, r)
	case http.MethodDelete:
		apiCfg.handlerDeleteUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := apiCfg.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)

	}
	log.Println(http.StatusCreated, user)

	respondWithJSON(w, http.StatusCreated, user)
}

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	if err := apiCfg.dbClient.DeleteUser(email); err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	log.Println(http.StatusOK, struct{}{})
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	user, err := apiCfg.dbClient.GetUser(email)
	if err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	log.Println(http.StatusOK, user)
	respondWithJSON(w, http.StatusOK, user)
}

func (apiCfg apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")

	var p struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := apiCfg.dbClient.UpdateUser(email, p.Password, p.Name, p.Age)
	if err != nil {
		log.Println(http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	log.Println(http.StatusOK, user)

	respondWithJSON(w, http.StatusOK, user)
}
