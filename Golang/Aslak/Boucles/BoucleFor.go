package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var age int

    for age < 18 { // on quitte la boucle s'il est mineur
        fmt.Print("Entrez votre age : ")
        scanner.Scan()
        age, _ = strconv.Atoi(scanner.Text())
    }

    fmt.Println("Bienvenue en boite de nuit !")
}

/*
Résultat :

Entrez votre age : 17
Entrez votre age : 19
Bienvenue en boite de nuit 
*/
