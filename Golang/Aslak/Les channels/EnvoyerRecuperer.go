/*
Pour envoyer ou recevoir une valeur dans un channel, il faut utiliser l'opérateur <-.

Exemple :
*/

package main

import "fmt"

func run(c chan string, name string) {
    c <- name // envoyer une valeur d'un channel
}

func main() {
    canal := make(chan string)
    go run(canal, "Clément")
    fmt.Println(<-canal) // récupérer une valeur d'un channel
}

/*
Résultat :

Clément
*/
