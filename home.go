package main

import (
    "net/http"
    "fmt"
)


func homeHandler(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "index.html", nil)
    fmt.Println("Home...")
}
