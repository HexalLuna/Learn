/*
Comme on peut accéder à l'adresse mémoire d'une variable depuis un pointeur
alors il est tout à fait réalisable de modifier la valeur qui se trouve dans cette adresse mémoire.
*/

package main

import "fmt"

func main() {
	var a int = 20
	var ap *int
	ap = &a

	fmt.Printf("Valeur de la variable a : %d\n", a)

	*ap = 30 // modification de la valeur de la variable "a" depuis un pointeur

	fmt.Printf("Valeur de la variable a : %d\n", a)

}

/*
Résultat :

Valeur de la variable a : 20
Valeur de la variable a : 30
*/
/*
Voici l'explication du code :

*ap = 30
Sur cette ligne de code, j'accède directement à l'adresse mémoire de la variable a grâce à mon pointeur ap , que je remplace par une nouvelle valeur (ici 30).
*/
