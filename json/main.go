package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type Article struct {
    Id      string `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Summary string `json:"summary"`
    Author  string `json:"author"`
}

func main() {
    http.HandleFunc("/api/hello", handleHello)
    http.HandleFunc("/company", handleCompany)
    http.HandleFunc("/Address", handleAddress)
    http.HandleFunc("/article", handleArticle)

    fmt.Println("Server running on port : 3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req.Method)
    fmt.Println(req.URL.Path)
    fmt.Fprintf(w, "<h1>Welcome to NetXD!</h1>")
}
func handleAddress(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req.Method)
    fmt.Println(req.URL.Path)
    fmt.Fprintf(w, "<h1>Karapakkam</h1>")
}
func handleCompany(w http.ResponseWriter, req *http.Request) {
    fmt.Println(req.Method)
    fmt.Println(req.URL.Path)
    fmt.Fprintf(w, "<h1>NetXd Software Private Limited</h1>")
}

func handleArticle(w http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
        reqBody, _ := io.ReadAll(req.Body)
        var post Article
        err := json.Unmarshal(reqBody, &post)

        post.Author = "John"
        if err != nil {
            fmt.Fprintf(w, err.Error())
        } else {
            json.NewEncoder(w).Encode(post)
        }
    } else {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Unable to process POST request")
    }
}