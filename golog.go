package main

import (
    "fmt"
    "net/http"
    "runtime"
)

func test() {
    fmt.Printf("hello Golog!\n")
}

func main() {
    for rule, funcname := range Urls {
        http.HandleFunc("/", runtime.FuncForPC(funcname))
    }
}