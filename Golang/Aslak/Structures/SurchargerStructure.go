package main

import "fmt"

func main() {

    // déclaration de la structure nommée Humain
    type Humain struct {
        nom string
        age int
    }

    hum := Humain{"Clément", 19} // surchargement des valeurs par défaut 
    fmt.Println(hum)
}

/*
Résultat :

{Clément 19}
*/
// Comme résultat nous avons les valeurs par défaut des attributs contenues dans la structure Humain
