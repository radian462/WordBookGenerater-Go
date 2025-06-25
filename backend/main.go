package main

import (
    "net/http"
    "log"
)

func main() {
    fs := http.FileServer(http.Dir("frontend/static"))
    http.Handle("/", fs)

    log.Println("Server started at :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}