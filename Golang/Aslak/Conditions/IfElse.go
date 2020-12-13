package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Entrez votre age : ")
    scanner.Scan()                   
    age, err := strconv.Atoi(scanner.Text())
    if err != nil {
        // Si la conversion n'a pas fonctionné alors on affiche un message d'erreur et on quitte le programme
        fmt.Println("On essaie de m'arnaquer ? allé Dehors ! Et la prochaine entrez un entier !")
        os.Exit(2) // on quitte le programmation
    }

    if age < 17 { // vérifier si l'utilisateur à au moins 18 ans
        fmt.Println("Sortez !")
    } else { // si ce n'est pas le cas alors on l'accepte pas
        fmt.Println("Entrez :)")
    }
}

/* Résultat :

Entrez votre age : 16
Sortez !

*/
