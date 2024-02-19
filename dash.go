package main

import (
    "net/http"
)

func dashHandler(w http.ResponseWriter, r *http.Request){
        tmpl.ExecuteTemplate(w, "main.html", nil);
        return
}
