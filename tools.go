package main

import(
   "time"
)

func momentaryMessage(msg string) string {

    currentMessage := make(chan string, 1);

    text := msg;
    currentMessage <- text;

    time.Sleep(2 * time.Second);
    return text;
}

func passValidate(pass string, hash string) int {

    var aux int = -1;
    passhash := md5Encode(pass);

    if(passhash == hash){
        aux = 0; // el password es correcto
        return aux;
    }else{
        aux = 1; // el password es incorrecto
        return aux;
    }
    return aux; // error
}
