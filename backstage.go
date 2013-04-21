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
        err := orm.Get(&cat, "id=?", cat_id)
        if err != nil {
            http.Redirect(w, r, "/admin/modify_cat", 301)
        } else {
            layout_data.Title = cat.Display_name
        }
    }
    if r.Method == "POST" {
        cat.Name = r.FormValue("name")
        cat.Display_name = r.FormValue("display_name")
        fmt.Printf("%v", cat)
        err := orm.Save(&cat)
        fmt.Printf("%v", err)
    } else {
        renderLayout(w, "modify_cat", layout_data, cat, "b_layout")
    }
    /*
    if r.Method == "POST" {
        cat_attributes := map[string]interface{} {
            "name" : r.FormValue("name"),
            "display_name" : r.FormValue("display_name"),
        }
        cat_id := cat_model.insert(cat_attributes)
        url := fmt.Sprintf("/admin/modify_cat?id=%d", cat_id)
        fmt.Printf("%s", url)
        http.Redirect(w, r, fmt.Sprintf("/admin/modify_cat?id=%d", cat_id), 301)
    } else {
        layout_data := layoutData{Title : "Create Category"}
        index_data := bIndexData{Title : "创建分类"}
        renderLayout(w, "create_cat", layout_data, index_data, "b_layout")
    }
    */
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