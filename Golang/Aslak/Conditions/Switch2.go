package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "time"
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

    switch {
    case choix >= 2000:
        println("Ah un 2000 !")
    case choix >= 1939 && choix <= 1945:
        println("Triste année")
    case time.Now().Weekday() == time.Sunday:
        println("Dimanche !")
    default:
        println("Mauvais choix !") // Valeur par défaut
    }
}

/*
------------------------------
Résultat:

Votre choix : 2000
Ah un 2000 !
------------------------------
Votre choix : 1942
Triste année
------------------------------
Votre choix : 0
Mauvais choix !
------------------------------
*/
