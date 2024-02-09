package main

import (
    "flag"
    "fmt"
    "net/http"
    "text/template"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("public/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    numCount := flag.Int("count", 10, "Number of values to generate")
    flag.Parse()

    data := struct {
        Numbers     []int
        Factorials  []int
        Fibonaccis  []int
    }{}

    for i := 1; i <= *numCount; i++ {
        data.Numbers = append(data.Numbers, i)
        data.Factorials = append(data.Factorials, factorial(i))
        data.Fibonaccis = append(data.Fibonaccis, fibonacci(i))
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
