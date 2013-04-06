package main

import (
    "net/http"
)

var Urls = map[string]func(http.ResponseWriter, *http.Request) {
    "/" : index,
    "/login" : login,
    "/admin/" : b_index,
    "/admin/comment" : b_comment,
}