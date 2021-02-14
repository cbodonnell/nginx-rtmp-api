package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func publish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "publish")
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
	fmt.Fprintln(w, "publish_done")
	vars := mux.Vars(r)
	fmt.Println("publish_done : " + r.RemoteAddr + " : " + vars["name"])

	time.Sleep(12 * time.Second)

	stream, err := stopStream(vars["name"])
	savedPath := fmt.Sprintf("/var/www/vod/%d/%d", stream.UserID, stream.ID)
	fmt.Println("Saving stream to " + savedPath)

	err = CopyDirectory("/var/www/hls", savedPath)
	if err != nil {
		internalServerError(w, err)
		return
	}
}

func clearStreams(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "admin/clear_streams")
	vars := mux.Vars(r)
	fmt.Println("admin/clear_streams : " + r.RemoteAddr + " : " + vars["user_id"])

	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	err = deleteStreams(userID)
	if err != nil {
		badRequest(w, err)
		return
	}

	err = RemoveContents("/var/www/vod/" + vars["user_id"])
	if err != nil {
		internalServerError(w, err)
		return
	}
}
