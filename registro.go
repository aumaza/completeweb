package main

import (
    "net/http"
    "fmt"
    "time"
    "log"
)

func registroHandler(w http.ResponseWriter, r *http.Request){

        db = connectionDB();
        defer db.Close();

        Nombre :=  r.PostFormValue("nombre_apellido");
        Email :=  r.PostFormValue("email");
        Password_1 := r.PostFormValue("pwd_1");
        Password_2 := r.PostFormValue("pwd_2");
        Celular := r.PostFormValue("celular");
        Direccion := r.PostFormValue("direccion");
        Localidad := r.PostFormValue("localidad");
        Provincia := r.PostFormValue("provincia");

        tmpl.ExecuteTemplate(w, "registro.html", nil);

        if r.Method != http.MethodPost {
            return;
        } else {

        var userExists bool

        err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM usuarios WHERE nombre = ? AND user = ?", Nombre, Email).Scan(&userExists);

        if err != nil {
            log.Fatal(err)
        }

        if(userExists == false) {

        now := time.Now();
        fmt.Println("=======================================");
        fmt.Println("Current date and time:", now);
        fmt.Println("=======================================");

        if (len(Nombre) == 0) ||
                (len(Email) == 0) ||
                    (len(Password_1) == 0) ||
                        (len(Password_2) == 0) ||
                            (len(Celular) == 0) ||
                                (len(Direccion) == 0) ||
                                    (len(Localidad) == 0) ||
                                        (len(Provincia) == 0) {

                                            fmt.Println("Hay Campos sin completar");
                                            fmt.Fprintf(w, momentaryMessage(message_empty_fields));
        } else {

            if (len(Password_1) <= 10) && (len(Password_2) <= 10) {

                    if (Password_1 == Password_2) {

                        hash := md5Encode(Password_1);


                        usr := Usuario {
                            nombre:Nombre,
                            user:Email,
                            email:Email,
                            password:hash,
                            celular:Celular,
                            direccion:Direccion,
                            localidad:Localidad,
                            provincia:Provincia,
                        }

                       insert ,err := db.Query("insert into usuarios (nombre, user, password, email, celular, direccion, localidad, provincia) values(?,?,?,?,?,?,?,?)",usr.nombre,usr.user,usr.password,usr.email,usr.celular,usr.direccion,usr.localidad,usr.provincia);

                       if err != nil{
                           fmt.Fprintf(w, "Warning. There was an error...");
                           fmt.Fprintf(w, momentaryMessage(err.Error()));
                        }else{
                            fmt.Println("Registro Exitoso");
                            fmt.Fprintf(w, message_wait);
                            fmt.Fprintf(w, momentaryMessage(message_ok));
                            defer insert.Close()
                        }

                    }else{
                        fmt.Println("Las contraseñas no coinciden");
                        fmt.Fprintf(w, momentaryMessage(message_password_unmatch));
                    }

                } else {
                    fmt.Println("Las contraseñas deben tener 10 caracteres como máximo");
                    fmt.Fprintf(w, momentaryMessage(message_password_length));
                }
        }

        }
        if(userExists == true){
            fmt.Println("Usuario Existente!. Pruebe con otros datos o verifique si olvidó su contraseña");
            fmt.Fprintf(w, momentaryMessage(message_user_exists));
        }

        fmt.Println("Nombre: ", Nombre);
        fmt.Println("Email: ", Email);
        fmt.Println("Celular: ", Celular);
        fmt.Println("Direccion: ", Direccion);
        fmt.Println("Localidad: ", Localidad);
        fmt.Println("Provincia: ", Provincia);

        fmt.Println("=======================================");

        }

}
