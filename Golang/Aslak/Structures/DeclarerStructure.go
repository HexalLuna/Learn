package main

import "fmt"

func main() {

    // déclaration de la structure nommée Humain
    type Humain struct {
        nom string
        age int
    }

    var hum Humain // initialisation d'une variable de type Humain
    fmt.Println(hum)
}

/*
Résultat :

{ 0}
*/
