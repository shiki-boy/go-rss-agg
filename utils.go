package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to marshal json: ", payload, err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	type errRes struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errRes{msg})
}
