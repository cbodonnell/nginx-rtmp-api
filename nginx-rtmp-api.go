package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/otiai10/copy"
)

func publish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "publish")
	vars := mux.Vars(r)
	fmt.Println("publish : " + r.RemoteAddr + " : " + vars["name"])

	err := RemoveContents("/var/www/hls")
	if err != nil {
		panic(err)
	}
}

func publishDone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "publish_done")
	vars := mux.Vars(r)
	fmt.Println("publish_done : " + r.RemoteAddr + " : " + vars["name"])

	// TODO: Get stream id to name folder
	id := 1

	savedPath := fmt.Sprintf("/var/www/vod/%s/%d", vars["name"], id)

	err := copy.Copy("/var/www/hls", savedPath)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/publish/{name}", publish)
	r.HandleFunc("/publish_done/{name}", publishDone)

	port := 9000
	fmt.Println(fmt.Sprintf("Serving on port %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
