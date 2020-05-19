package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    fmt.Println("Code.education Rocks!")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Code.Education Rocks! \nTente a url /ola/seu_nome")
    })

    http.HandleFunc("/ola/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/ola/"):]
        fmt.Fprintf(w, "Olá %s\nCode.education Rocks!", name)
    })

    log.Println("** Iniciado na porta: " + port + " **")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }

}