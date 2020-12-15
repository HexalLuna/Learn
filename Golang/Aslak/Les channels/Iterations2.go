/*
Pour nous prémunir de cette erreur, il faut indiquer à notre compilateur qu'il n'est pas nécessaire de lire la suite du buffer de notre channel,
cela est possible avec la fonction close().
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 2) // buffer de taille 2

	go func() {
		defer close(ch) // on indique à notre compilateur qu'on a finit d'écrire sur le channel
		ch <- "test"
	}()

	for elem := range ch {
		fmt.Println(elem)
	}

}

/*
Résultat :

test
*/
