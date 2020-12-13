// Fonction sans type de retour mais avec des paramètres
// Une fonction est capable de prendre autant de paramètres que vous voulez et peu importe leurs types.

package main

import (
    "fmt"
)

// prend en paramètre un type string et un int
func affichage(nom string, age int) {
    fmt.Println("Bonjour", nom, "vous avez", age, "ans")
}

func main() {
    affichage("Clément", 8)
    affichage("Alexis", 16)
}

/*
Résultat :

Bonjour Clément vous avez 8 ans
Bonjour Alex vous avez 16 ans
*/
