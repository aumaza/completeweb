package main

import (
 "net/http"
 "fmt"
 "log"
)

func passHandler(w http.ResponseWriter, r *http.Request){

    db = connectionDB();
    defer db.Close();


    user :=  r.PostFormValue("user");
    Password_1 := r.PostFormValue("pwd_1");
    Password_2 := r.PostFormValue("pwd_2");

    tmpl.ExecuteTemplate(w, "password.html", nil);

    if r.Method != http.MethodPost {
            return;
    }

        if((len(user) == 0) ||
                (len(Password_1) == 0) ||
                    (len(Password_2) == 0)) {

                        fmt.Println("Atención!Hay campos sin completar");
                        fmt.Fprintf(w, momentaryMessage(message_pass_empty_fields));

                    }else{

                            var userExists bool

                            err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM usuarios WHERE user = ?",  user).Scan(&userExists);

                            if err != nil {
                                log.Fatal(err)
                            }

                            if (userExists == false) {

                                fmt.Println("Usuario inexistente. Registresé previamente");
                                fmt.Fprintf(w, momentaryMessage(message_pass_user_no_exists));

                            }

                            if(userExists == true) {

                                 if ((len(Password_1) <= 10) && (len(Password_2) <= 10)){

                                     if (Password_1 == Password_2){

                                         passhash := md5Encode(Password_1);

                                         update ,err := db.Query("update usuarios set password = ? WHERE user = ?", passhash, user);

                                         if err != nil {
                                            fmt.Fprintf(w, "Warning. There was an error...");
                                            fmt.Fprintf(w, momentaryMessage(err.Error()));
                                            }else{
                                                fmt.Println("Se actualizó la contraseña exitosamente");
                                                fmt.Fprintf(w, message_wait);
                                                fmt.Fprintf(w, momentaryMessage(message_pass_ok));
                                                defer update.Close()
                                            }

                                    }else{

                                            fmt.Println("Atención. Las contraseñan no coinciden");
                                            fmt.Fprintf(w, momentaryMessage(message_pass_unmatch));
                                    }

                                }else{

                                            fmt.Println("Las contraseñas deben tener 10 caracteres como máximo");
                                            fmt.Fprintf(w, momentaryMessage(message_pass_lengh));
                                }
                            }
                    }

}
