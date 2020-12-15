/*
il y a deux manières pour écrire dans un fichier. 
Soit vous décidez d'écraser un fichier après écriture, soit d'écrire à la suite du contenu du fichier.

Pour écrire dans un fichier on va utiliser la bibliothèque os.
*/

package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    defer file.Close() // on ferme automatiquement à la fin de notre programme

    if err != nil {
        panic(err)
    }

    _, err = file.WriteString("test\n") // écrire dans le fichier
    if err != nil {
        panic(err)
    }

    _, err = file.WriteString("i love test\n")
    if err != nil {
        panic(err)
    }

    data, err := ioutil.ReadFile("test.txt") // lire le fichier
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))
}

*/
Information

J'utilise le "\n" pour faire un saut de la ligne dans mon fichier.
-------------------------------
Résultat :

je suis un fichier de test
test
i love test
-------------------------------
*/
