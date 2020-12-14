// La valeur par défaut d'un pointeur est nil. Ça peut vous aider pour savoir si votre pointeur comporte une adresse mémoire ou pas, exemple :

package main

import "fmt"

func main() {
    var ptr *int

    if ptr == nil { // est ce que le pointeur est vide ?
        fmt.Println("Aucune adresse mémoire")
    } else {
        fmt.Printf("Votre adresse mémoire est %p\n", ptr)
    }

    var b int = 20
    var pb *int
    pb = &b

    if pb == nil {
        fmt.Println("Aucune adresse mémoire")
    } else {
        fmt.Printf("Votre adresse mémoire est %p\n", pb)
    }
}

/*
Résultat :

Aucune adresse mémoire
Votre adresse mémoire est 0xc000056058
*/
