package main

import (
 "net/http"
 "fmt"
)

func passHandler(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "password.html", nil)
    fmt.Println("Going to change Password...")
}
