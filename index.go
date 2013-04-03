package main

import (
    "net/http"
)

type indexData struct {
    Content string
}

func index(w http.ResponseWriter, r *http.Request) {
    //render(w, "index")
    
    index_data := indexData{Content : "这是我要显示在index模板的一句话"}
    layout_data := layoutData{Title : "Golang Blog"}
    renderLayout(w, "index", layout_data, index_data)
}

func login(w http.ResponseWriter, r *http.Request) {
    renderSingle(w, "login")
}