package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func publish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "publish")
	vars := mux.Vars(r)
	fmt.Println("publish : " + r.RemoteAddr + " : " + vars["name"])
	os.RemoveAll(fmt.Sprintf("/var/www/hls/%s.m3u8", vars["name"]))
	os.RemoveAll(fmt.Sprintf("/var/www/hls/%s_low/", vars["name"]))
	os.RemoveAll(fmt.Sprintf("/var/www/hls/%s_mid/", vars["name"]))
	os.RemoveAll(fmt.Sprintf("/var/www/hls/%s_hi/", vars["name"]))
}

func publishDone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "publish_done")
	vars := mux.Vars(r)
	fmt.Println("publish_done : " + r.RemoteAddr + " : " + vars["name"])

	writeLineToFile(fmt.Sprintf("/var/www/hls/%s/index.m3u8", vars["name"]), "#EXT-X-ENDLIST")
	// writeLineToFile(fmt.Sprintf("/var/www/hls/%s_mid/index.m3u8", vars["name"]), "#EXT-X-ENDLIST")
	// writeLineToFile(fmt.Sprintf("/var/www/hls/%s_hi/index.m3u8", vars["name"]), "#EXT-X-ENDLIST")
}

func writeLineToFile(filename string, text string) {
	fmt.Println(fmt.Sprintf("Writing %s to %s", text, filename))
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/publish/{name}", publish)
	r.HandleFunc("/publish_done/{name}", publishDone)

	log.Fatal(http.ListenAndServe(":9000", r))
}
