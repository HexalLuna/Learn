package main //Merci à Aslak de m'avoir expliquer pourquoi utiliser le package bufio pour ce type d'entrée utilisateur et non le package fmt

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
