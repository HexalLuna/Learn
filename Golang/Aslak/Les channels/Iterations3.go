/*
Il existe aussi un autre moyen d'itérer sur un channel, mais avant de vous montrez le code 
il faut savoir qu'un channel renvoie à la fois sa valeur mais aussi un booléen qui vous indique si la valeur a été envoyée sur le channel.
*/

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2) // buffer de taille 2

	go func() {
		defer close(ch) // on indique à notre compilateur qu'on a finit d'écrire sur le channel
		ch <- "test"
	}()

	for true {
		if elem, ok := <-ch; ok == true { // est ce que le chanel possède encore un récepteur ?
			fmt.Println(elem, ok)
		} else {
			break
		}
	}
}

/*
Résultat :

test
*/
