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
/**
 * bufio.NewScanner() encapsule un objet io.Reader ou io.Writer, créant un autre objet (Reader ou Writer) qui implémente également l'interface
 * en d'autres termes, tout ce qu'il fait, c'est ajouter une couche par exemple dans os.Stdin, il n'effectue aucune analyse ou interprétation du flux.
 *
 * En revanche, fmt.Scanln() lit les données à partir d'un flux (qui peut ou non être mis en mémoire tampon) 
 * c'est-à-dire renvoyé par le package bufio, en divisant l'entrée par un espace, pour la stocker dans une slice.
**/ 
