#!/bin/bash
clear
rm mybills
echo "==================================================================================="
echo "Compilando ejecutable para GNU/Linux..."
echo "==================================================================================="
if [[ ! -f mybills ]]
then
GOOS=linux GOARCH=amd64 go build -o mybills main.go home.go registro.go password.go connection.go usuarios.go tools.go registro_messages.go password_messages.go
echo "Ejecutable creado con éxito..."
echo "==================================================================================="
    echo "Desea ejecutar el programa? (s/n)"
    read resp
    if [[ $resp == s ]]
    then
    ./mybills
    elif [[ $resp == n ]]
    then
    echo "==================================================================================="
    echo "Solo se compiló el binario"
    echo "==================================================================================="
    else
    echo "==================================================================================="
    echo "Debe responder con s | n"
    echo "==================================================================================="
    fi
else
echo "ERROR, Archivo existente..."
echo "==================================================================================="
fi
#rm first.exe
#echo "==================================================================================="
#echo "Compilando ejecutable para MS Windows..."
#echo "==================================================================================="
#GOOS=windows GOARCH=amd64 go build -fPIE -o first.exe main.go show_time.go lib.go employers.go
