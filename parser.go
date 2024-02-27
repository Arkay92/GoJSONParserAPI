package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Person struct {
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    City    string   `json:"city"`
    Friends []Friend `json:"friends"`
}

type Friend struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func personHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
        return
    }

    var person Person
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&person)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Received: %+v\n", person)
}

func main() {
    http.HandleFunc("/person", personHandler)

    fmt.Println("Server is listening on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}
