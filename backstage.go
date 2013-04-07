/* backstage.go */
package main

import (
    "net/http"
    "fmt"
)

type bIndexData struct {
    Title string
}

func b_index(w http.ResponseWriter, r *http.Request) {
    layout_data := layoutData{Title : "Golang Blog"}
    index_data := bIndexData{Title : "backstage"}
    renderLayout(w, "b_index", layout_data, index_data, "b_layout")
}

func create_post(w http.ResponseWriter, r *http.Request) {

}

func modify_cat(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        name := r.FormValue("name")
        display_name :=  r.FormValue("display_name")
        fmt.Fprintf(w, "name: %s, display_name: %s", name, display_name)
        //http.Redirect(w, r, "/admin/modify_cat?id=")
    } else {
        layout_data := layoutData{Title : "Create Category"}
        index_data := bIndexData{Title : "创建分类"}
        renderLayout(w, "create_cat", layout_data, index_data, "b_layout")
    }
}

func create_page(w http.ResponseWriter, r *http.Request) {

}

func b_page(w http.ResponseWriter, r *http.Request) {

}

func b_cat(w http.ResponseWriter, r *http.Request) {

}

func b_comment(w http.ResponseWriter, r *http.Request) {
    layout_data := layoutData{Title : "Golang Blog"}
    index_data := bIndexData{Title : "backstage"}
    renderLayout(w, "b_comment", layout_data, index_data, "b_layout")
}