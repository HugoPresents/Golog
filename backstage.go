/* backstage.go */
package main

import (
    "net/http"
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

func create_cat(w http.ResponseWriter, r *http.Request) {

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