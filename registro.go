package main

import (
    "net/http"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "time"
)

func registroHandler(w http.ResponseWriter, r *http.Request){
    tmpl.ExecuteTemplate(w, "registro.html", nil)

    if r.Method != http.MethodPost {
        return
    }
        now := time.Now()
        fmt.Println("=======================================")
        fmt.Println("Current date and time:", now)
        fmt.Println("Se presionó bonton submit...")
        fmt.Println("=======================================")

        //r.ParseForm()

        Nombre :=  r.PostFormValue("nombre_apellido")
        Email :=  r.PostFormValue("email")
        Password := r.PostFormValue("pwd_1")
        Celular := r.PostFormValue("celular")
        Direccion := r.PostFormValue("direccion")
        Localidad := r.PostFormValue("localidad")
        Provincia := r.PostFormValue("provincia")

        fmt.Println("Nombre: ", Nombre)
        fmt.Println("Email: ", Email)
        fmt.Println("Password: ", Password)
        fmt.Println("Celular: ", Celular)
        fmt.Println("Direccion: ", Direccion)
        fmt.Println("Localidad: ", Localidad)
        fmt.Println("Provincia: ", Provincia)
        fmt.Println("=======================================")

    /*
     if err := r.ParseForm(); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
     }
    */
}


func md5Encode(input string) string {

	// Create a new hash & write input string
	hash := md5.New()
	_, _ = hash.Write([]byte(input))

	// Get the resulting encoded byte slice
	md5 := hash.Sum(nil)
    passHash := hex.EncodeToString(md5[:])

	// Convert the encoded byte slice to a string
    return passHash
	//return fmt.Sprintf("%x", md5)
}
