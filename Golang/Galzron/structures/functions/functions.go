//Ici nous allons voir comment utiliser une fonction dans une autre
package main 

import "fmt"

type entier int

func main()  {

	var e int = 3
	var r int = 2

	if somme(e, r) == 5 {
		fmt.Println("Les valeurs de e et r sont corrects")
	} else {
		fmt.Println("Les valeurs de e et r sont incorrects")
	}
}

func somme(a, b int) entier {
	return entier(a + b)
}