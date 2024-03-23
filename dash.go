package main

import (
    "net/http"
)


func dashHandler(w http.ResponseWriter, r *http.Request){
    session, _ := store.Get(r, "session-name")

    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    tmpl.ExecuteTemplate(w, "main.html", nil);
        return
}
