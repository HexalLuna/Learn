/*
Il existe deux façons pour créer une Slice.

Soit à partir de la même syntaxe qu'un tableau sauf que cette fois-ci on ne spécifie pas la taille du tableau :
*/

package main

import (
    "fmt"
)

func main() {
    var nombres = []int{0, 0, 0, 0, 0} // création d'une slice avec 5 éléments 
    fmt.Println(nombres)
}

/*
Résultat :

[0 0 0 0 0]
*/

// Soit depuis la fonction make()

package main

import (
    "fmt"
)

func main() {
    var nombres = make([]int, 5) // création d'une slice avec 5 éléments
    fmt.Println(nombres)
}

/*
Résultat :

[0 0 0 0 0]
*/
