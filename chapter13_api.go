// Chapter 13: Building RESTful APIs in Go
package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Article struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

var articles = []Article{}

func listArticles(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
    var a Article
    if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }
    articles = append(articles, a)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(a)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/v1/articles", listArticles).Methods("GET")
    r.HandleFunc("/api/v1/articles", createArticle).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", r))
}
