package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/plus3it/gorecurcopy"
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

	savedPath := fmt.Sprintf("/var/www/hls/%s/%d", vars["name"], id)

	if _, err := os.Stat(savedPath); os.IsNotExist(err) {
		os.Mkdir(savedPath, os.ModeDir)
	}
	err := gorecurcopy.CopyDirectory("/var/www/hls", savedPath)
	if err != nil {
		panic(err)
	}
}

// func writeLineToFile(filename string, text string) {
// 	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	if _, err = f.WriteString(text); err != nil {
// 		panic(err)
// 	}
// }

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/publish/{name}", publish)
	r.HandleFunc("/publish_done/{name}", publishDone)

	log.Fatal(http.ListenAndServe(":9000", r))
}
