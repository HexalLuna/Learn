package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    fmt.Print("Entrez quelque chose, n'importe quoi fera l'affaire puis appuyez sur entrer : ")
    bufio.NewScanner(os.Stdin).Scan() // création du scanner capturant l'entrée utilisateur et lancement du scanner
    
    entreeUtilisateur := bufio.NewScanner(os.Stdin).Text() // stockage du résultat du scanner dans une variable
    fmt.Println(entreeUtilisateur)
}
