package main

import (
    "fmt"
    "net/http"
    "text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("public/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    numbers := []int{}
    for i := 1; i <= 10; i++ {
        numbers = append(numbers, i)
    }

    data := struct {
        Numbers []int
    }{
        Numbers: numbers,
    }

    tmpl.Execute(w, data)
}

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handler)
    fmt.Println("Server listening on port 8080")
    http.ListenAndServe(":8080", nil)
}
