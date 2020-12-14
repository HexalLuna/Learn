// On utilisera la fonction Printf() avec le symbole %p et le signe & pour accéder et afficher à l'adresse d'une variable

package main

import (
    "fmt"
)

func main() {
    var a int = 20
    fmt.Printf("Adresse de la variable a: %p\n", &a)
}

/*
Résultat :

Adresse de la variable a: 0xc0000100a0
*/
