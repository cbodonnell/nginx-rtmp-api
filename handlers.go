package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func publish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "publish")
	vars := mux.Vars(r)
	fmt.Println("publish : " + r.RemoteAddr + " : " + vars["name"])

	_, err := startStream(vars["name"])
	if err != nil {
		badRequest(w, err)
		return
	}

	err = RemoveContents("/var/www/hls")
	if err != nil {
		internalServerError(w, err)
		return
	}
}

func publishDone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "publish_done")
	vars := mux.Vars(r)
	fmt.Println("publish_done : " + r.RemoteAddr + " : " + vars["name"])

	id, err := stopStream(vars["name"])
	savedPath := fmt.Sprintf("/var/www/vod/%s/%d", vars["name"], id)
	fmt.Println("Saving stream to ", savedPath)

	err = CopyDirectory("/var/www/hls", savedPath)
	if err != nil {
		internalServerError(w, err)
		return
	}
}

// TODO: Move to a separate public api w/ auth
func getStream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("get_stream : " + r.RemoteAddr + " : " + vars["user_id"])

	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	stream, err := queryStream(userID)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stream)
}

func getStreams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("get_streams : " + r.RemoteAddr + " : " + vars["user_id"])

	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	streams, err := queryStreams(userID)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(streams)
}
