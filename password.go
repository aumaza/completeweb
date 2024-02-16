package main

import (
 "net/http"
 "fmt"
 "time"
 "log"
)

func passHandler(w http.ResponseWriter, r *http.Request){

    db = connectionDB();
    defer db.Close();

    Email :=  r.PostFormValue("email");
    Password_1 := r.PostFormValue("pwd_1");
    Password_2 := r.PostFormValue("pwd_2");

    tmpl.ExecuteTemplate(w, "password.html", nil);

    if r.Method != http.MethodPost {
            return;
    }

}
