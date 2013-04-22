/* backstage.go */
package main

import (
    "net/http"
    "fmt"
    "strconv"
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
    id := r.FormValue("id")
    var cat Category
    layout_data := layoutData{Title : "Create Category"}
    if id != "" {
        cat_id, _ := strconv.Atoi(id)
        err := orm.Where("id=?", cat_id).Find(&cat)
        checkErr(err)
        if err != nil {
            http.Redirect(w, r, "/admin/modify_cat", 301)
        } else {
            layout_data.Title = cat.DisplayName
        }
    }
    if r.Method == "POST" {
        cat.Name = r.FormValue("name")
        cat.DisplayName = r.FormValue("display_name")
        orm.Save(&cat)
        url := fmt.Sprintf("/admin/modify_cat?id=%d", cat.Id)
        http.Redirect(w, r, url, 301)
    } else {
        renderLayout(w, "modify_cat", layout_data, cat, "b_layout")
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