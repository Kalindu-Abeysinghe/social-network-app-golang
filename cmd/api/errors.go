package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJsonError(w, http.StatusInternalServerError, "Internal server error")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResposne(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Found error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJsonError(w, http.StatusNotFound, "Not found")
}