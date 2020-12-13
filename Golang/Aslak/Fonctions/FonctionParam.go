package main

import (
    "fmt"
)

func main() {
    fmt.Println(addition(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13))
}

// déclaration d'une fonction avec des paramètres infinis
func addition(param ...int) int {
    total := 0
    for _, value := range param { //j'ai mis un underscore "_" car je ne souhaite pas récupérer l'index de la range
        total += value
    }
    return total
}

/*
Résultat :

91
*/
