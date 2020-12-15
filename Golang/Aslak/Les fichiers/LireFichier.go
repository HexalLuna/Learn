/*
Il existe de multiples façons pour lire et écrire dans un fichier. 

Je vais manipuler un fichier nommé test.txt avec le contenu suivant :

je suis un fichier de test


Dans le cas ou vous ne faites que lire un fichier, le mieux reste d'utiliser la fonction ReadFile() de la bibliothèque io/ioutil. 
Cette fonction retourne un type byte (bit), il faut donc penser à caster (convertir le type) le résultat obtenu en string
*/

package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("test.txt") // lire le fichier text.txt
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(data)) // conversion de byte en string
}

/*
---------------------------------
Résultat :

je suis un fichier de test
---------------------------------
Dans le cas ou le fichier n'existe pas, vous aurez l'erreur suivante :

Le fichier spécifié est introuvable.
*/
