package main

import (
	"fmt"
	"net/http"
)

func badRequest(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
	var msg string
	if config.Debug {
		msg = err.Error()
	} else {
		msg = "Bad request"
	}
	http.Error(w, msg, http.StatusBadRequest)
}

func internalServerError(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
	var msg string
	if config.Debug {
		msg = err.Error()
	} else {
		msg = "Internal server error"
	}
	http.Error(w, msg, http.StatusInternalServerError)
}
