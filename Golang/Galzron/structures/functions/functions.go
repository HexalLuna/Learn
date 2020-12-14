//Ici nous allons voir comment utiliser une fonction dans une autre, pour rendre cela plus jolie je vais utiliser les entrées utilisateurs
package main 

import "fmt"

type entier int

func main()  {

	var e int
	var r int

	fmt.Println("Ecrivez ci dessous deux valeurs, qui une fois additioner donner le chiffre 5")
	fmt.Scanln(&e, &r)

	if somme(e, r) == 5 {
		fmt.Printf("Les valeurs de %d et %d sont corrects", e, r)
	} else {
		fmt.Printf("Les valeurs de %d et %d sont incorrects", e, r)
	}
}

func somme(a, b int) entier {
	return entier(a + b)
}
