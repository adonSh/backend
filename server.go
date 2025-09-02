package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

type Backend struct {
    *http.ServeMux
}

func main() {
    b := Backend { http.NewServeMux() }
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    b.HandleFunc("GET /", b.defaultHandler)

    log.Fatal(http.ListenAndServe("0.0.0.0:" + port, b))
}

func (b *Backend) defaultHandler(w http.ResponseWriter, r *http.Request) {
    timestamp := time.Now().UTC().Add(5 * time.Minute)
    fmt.Fprintf(w, "%d\n", timestamp.Unix())
}
