package main

import (
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
	hash := md5.New();
	_, _ = hash.Write([]byte(input));

	// Get the resulting encoded byte slice
	md5 := hash.Sum(nil);
    passHash := hex.EncodeToString(md5[:]);

	// Convert the encoded byte slice to a string
    return passHash;
	//return fmt.Sprintf("%x", md5)
}
