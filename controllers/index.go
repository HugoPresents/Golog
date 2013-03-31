package main

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello Golog!</h1>")
}