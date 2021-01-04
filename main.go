package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func publish(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "publish")
    vars := mux.Vars(r)
    fmt.Println("publish : " + r.RemoteAddr + " : " + vars["name"])
}

func publishDone(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "publish_done")
    vars := mux.Vars(r)
    fmt.Println("publish_done : " + r.RemoteAddr + " : " + vars["name"])
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/publish/{name}", publish)
    r.HandleFunc("/publish_done/{name}", publishDone)

    log.Fatal(http.ListenAndServe(":9000", r))
}
