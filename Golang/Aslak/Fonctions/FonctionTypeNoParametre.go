// Fonction sans type de retour et sans paramètres

package main

import (
    "fmt"
)

// déclaration de la fonction affichage()
func affichage() {
    fmt.Println("#################################")
    fmt.Println("\tBonjour")
    fmt.Println("#################################")
}

func main() {
    affichage() // appel de la fonction affichage()
}

/*
Résultat :

#################################
        Bonjour
#################################
*/
