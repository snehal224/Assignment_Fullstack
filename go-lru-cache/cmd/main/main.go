package main

import (
    "log"
    "net/http"
    "go-lru-cache/internal/handlers"
)

func main() {
    http.HandleFunc("/set", handlers.SetHandler)
    http.HandleFunc("/get", handlers.GetHandler)
    http.HandleFunc("/delete", handlers.DeleteHandler)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}
