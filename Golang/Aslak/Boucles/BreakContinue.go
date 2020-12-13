package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

func main() {

    rand.Seed(time.Now().UnixNano())
    randomInt := rand.Intn(10)

    scanner := bufio.NewScanner(os.Stdin)

    max := 20

    for true { // boucle infinie
        fmt.Print("Entrez votre nombre : ")
        scanner.Scan()
        nbr, err := strconv.Atoi(scanner.Text())

        if err != nil {
            fmt.Println("Entrez un nombre !")
            continue // on revient au début de la boucle
        }
        if nbr > max || nbr < 0 {
            fmt.Println("Votre nombre doit être compris entre 0 et ", max, " !")
            continue
        } else if nbr == randomInt {
            fmt.Println("Bien joué !")
            break // on quitte la boucle
        } else {
            fmt.Println("Dommage !")
        }
    }
}

/*
Résultat :

Entrez votre nombre : 50
Votre nombre doit être compris entre 0 et  20  !
------------------------------------------------
Entrez votre nombre : arnaque
Entrez un nombre !
------------------------------------------------
Entrez votre nombre : 6
Dommage !
------------------------------------------------
Entrez votre nombre : 15
Dommage !
------------------------------------------------
Entrez votre nombre : 3
Bien joué !
------------------------------------------------
*/
