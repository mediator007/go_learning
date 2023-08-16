package main

import (
    "fmt"
    "log"
    "net/http"
)

func handlerForEvent(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Event send!!!   ")
	go sender(8080)
}

func server() {

    http.HandleFunc("/", handlerForEvent)

    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Here is our API")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
