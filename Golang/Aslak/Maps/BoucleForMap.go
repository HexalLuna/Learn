// On peut utiliser la boucle for avec le mot-clé range pour récupérer la clé de tous les éléments de votre Map.

package main

import "fmt"

func main() {
    notes := map[string]int{"Clément": 20, "Alex": 18, "Kevin": 15, "Robert": 17}

	for eleve := range notes {
		fmt.Println("La note de ", eleve, "est", notes[eleve])
	}
}

/*
Résultat :

La note de  Clément est 20
La note de  Alex est 18
La note de  Kevin est 15
La note de  Robert est 17
*/
