package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, struct{}{})
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Something went wrong")
}
