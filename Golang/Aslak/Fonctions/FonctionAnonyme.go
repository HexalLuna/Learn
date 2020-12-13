package main

import (
    "fmt"
    "math"
)

// Déclaration d'une fonction qui prend en paramètres un float64 et une fonction anonyme
func rajouterDix(a float64, fAnonyme func(float64) float64) /**/ {
    operation := fAnonyme(a) // Appel à notre fonction anonyme
    result := operation + 10
    fmt.Println(result)
}

func main() {
    // stockage de notre fonction anonyme dans une variable
    racineCarree := func(x float64) float64 { return math.Sqrt(x) }
    rajouterDix(9, racineCarree)

    /* 
        il est possible aussi d'utiliser directement une fonction anonyme 
        dans une variable sans forcement la stocker dans une variable
    */
    rajouterDix(5, func(x float64) float64 { return math.Pow(x, 2) })
}

/*
Résultat :

13
35
*/
