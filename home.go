package main

import (
    "net/http"
    "fmt"
    "log"
)


func homeHandler(w http.ResponseWriter, r *http.Request){

    db = connectionDB();
    defer db.Close();

    user :=  r.PostFormValue("user");
    password := r.PostFormValue("pwd");

    tmpl.ExecuteTemplate(w, "index.html", nil);

    if r.Method != http.MethodPost {
            return;
    }

        if((len(user) == 0) || (len(password) == 0)) {
            fmt.Println("Debe ingresar sus datos...");
        } else {
                            var userExists bool

                            err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM usuarios WHERE user = ?",  user).Scan(&userExists);

                            if err != nil {
                                log.Fatal(err)
                                return
                            }

                            if (userExists == false) {

                                fmt.Println("Usuario inexistente. Registresé previamente o verifique el mail ingresado");
                                fmt.Fprintf(w, momentaryMessage(message_pass_user_no_exists));

                            }

                            if(userExists == true) {
                                var hash string;

                                er := db.QueryRow("SELECT password FROM usuarios WHERE user = ?", user).Scan(&hash);

                                var resp int = passValidate(password, hash);

                                if er != nil {
                                    log.Fatal(err)
                                    return
                                }

                                if(resp == -1){
                                    fmt.Println("Error...");
                                }
                                if(resp == 0){
                                    fmt.Println("Bienvenido/a al dashboard de My Bills"); // voy a la url main
                                    http.Redirect(w, r, "/dash", http.StatusMovedPermanently)
                                }
                                if(resp == 1){
                                    fmt.Println("Contaseña erronea. Intente nuevamente!!");
                                }


                            }

        }
}
