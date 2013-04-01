package main

import (
    "net/http"
)

var Urls = map[string]func(http.ResponseWriter, *http.Request) {
    "/" : index,
}