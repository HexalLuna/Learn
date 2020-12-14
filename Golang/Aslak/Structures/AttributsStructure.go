// Pour accéder et modifier votre attribut de votre structure, nous utiliserons un opérateur d’accès . (un point).

package main

import "fmt"

func main() {
    type Humain struct {
        nom string
        age int
    }

    hum := Humain{"Clément", 20}
    fmt.Println(hum)

    // accès juste à l'attribut age de notre structure Humain
    fmt.Println("je ne veux afficher que l'age de l'humain => ", hum.age)

    hum.age = 23 // modification de l'age
    fmt.Println("3 ans plus tard ...", hum.age)
}

/*
Résultat :

{Clément 20]
je ne veux afficher que l'age de l'humain =>  20
3 ans plus tard ... 23
*/
