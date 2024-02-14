package main

import (
    "net/http"
    "crypto/md5"
    "encoding/hex"
)


type Usuario struct {

    id int
    nombre string
    user string
    password string
    email string
    celular string
    direccion string
    localidad string
    provincia string

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

func insertUser(r *http.Request) int {

    var aux int = -1

    if r.FormValue("nombre") != "" || r.FormValue("email") != "" || r.FormValue("pwd_1") != "" || r.FormValue("pwd_2") != "" || r.FormValue("celular") != "" || r.FormValue("direccion") != "" || r.FormValue("localidad") != "" || r.FormValue("provincia") != "" {

        if (len([]rune(r.FormValue("pwd_1"))) == 10) && (len([]rune(r.FormValue("pwd_2"))) == 10) {

                if (r.FormValue("pwd_1") == r.FormValue("pwd_2")) {

                   /* usr := Usuario{
                        nombre: r.FormValue("nombre"),
                        user:r.FormValue("email"),
                        password:md5Encode(r.FormValue("pwd_1")),
                        email:r.FormValue("email"),
                        celular:r.FormValue("celular"),
                        direccion:r.FormValue("direccion"),
                        localidad:r.FormValue("localidad"),
                        provincia:r.FormValue("provincia"),
                    }*/
                   aux = 0
                   return aux
            } else{
                aux = 3 // las contraseñas no coinciden
                return aux
            }
     } else {
         aux = 2 // las contraseñas deben tener 10 caracteres como máximo
         return aux
    }
    } else {
        aux = 1 // hay campos vacios
        return aux
    }
        return aux
}
