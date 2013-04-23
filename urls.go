package main

import (
    "net/http"
)

var Urls = map[string]func(http.ResponseWriter, *http.Request) {
    "/" : index,
    "/login" : login,
    "/admin/" : bIndex,
    "/admin/comment" : bComment,
    "/admin/modify_cat" : modifyCat,
}