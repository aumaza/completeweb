package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/gorilla/sessions"
)

var (
    // Llave de autenticación. ¡Cámbialo por algo único y seguro!
    authKey = []byte("my_bills794613316497")
    store   = sessions.NewCookieStore(authKey)
)

func homeHandler(w http.ResponseWriter, r *http.Request){

    db = connectionDB();
    defer db.Close();

    session, _ := store.Get(r, "session-name")
    user :=  r.PostFormValue("user");
    password := r.PostFormValue("pwd");

   // tmpl.ExecuteTemplate(w, "index.html", nil);

    if r.Method == http.MethodPost {

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
                                        session.Values["authenticated"] = true
                                        session.Save(r, w)
                                        fmt.Println("Bienvenido/a al dashboard de My Bills"); // voy a la url main
                                        http.Redirect(w, r, "/dash", http.StatusMovedPermanently)
                                }
                                if(resp == 1){
                                    fmt.Println("Contaseña erronea. Intente nuevamente!!");
                                }


                            }

        }
    }


        tmpl.ExecuteTemplate(w, "index.html", nil);
}
