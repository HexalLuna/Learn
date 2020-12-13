package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Votre choix : ")
    scanner.Scan()
    choix, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println("Entrez un entier !")
        os.Exit(2)
    }

    switch choix { // la variable qu'on souhaite vérifier
        case 0, 1: // 1 ou 0
            println("Hamburger !")
        case 7:
            println("Entrecôte !")
        case 23:
            println("23 ans c'est l'âge où notre vie chnage le plus !")
        case 42:
            println("la réponse à la question ultime du sens de la vie!")
        case 666:
            println("Quand le diable veut une âme, le mal devient séduisant.")
        default:
            println("Mauvais choix !") // Valeur par défaut
    }
}

/*
------------------------------------------------------------
Résultat :

Votre choix : 666
Quand le diable veut une âme, le mal devient séduisant
------------------------------------------------------------
Votre choix : 7
Entrecôte !
------------------------------------------------------------
Votre choix : 1
Hamburger !
------------------------------------------------------------
Votre choix : 0
Hamburger !
------------------------------------------------------------
Votre choix : 50
Mauvais choix !
------------------------------------------------------------
*/
