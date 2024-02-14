package main

import (
    "fmt"
    "html/template"
    "net/http"
)

var tmpl *template.Template

func init(){
    tmpl = template.Must(template.ParseGlob("templates/*.html"))
}


func main () {

        db = connectionDB()

        if db != nil {
            fmt.Println("Database connection Successfully!!")
        } else {
            fmt.Println("Database connection Failure!")
        }

        fmt.Println("Server on: http://localhost:3000")
        fs := http.FileServer(http.Dir("assets"))
        http.Handle("/assets/",  http.StripPrefix("/assets", fs))
        http.HandleFunc("/", homeHandler)
        http.HandleFunc("/password", passHandler)
        http.HandleFunc("/registrarse", registroHandler)
        http.ListenAndServe(":3000", nil)

}
