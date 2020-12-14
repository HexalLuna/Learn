/*
Pour supprimer un élément de votre Map il faut utiliser la fonction remove(), qui prend comme paramètres d'abord votre Map et 
ensuite la clé de l'élément que vous voulez supprimer.
*/

package main

import "fmt"

func main() {
    notes := map[string]int{"Clément": 20, "Alex": 18, "Kevin": 15, "Robert": 17}
	fmt.Println(notes)

	delete(notes, "Clément")
	fmt.Println(notes)
}

/*
Résultat :

map[Alex:18 Clément:20 Kevin:15 Robert:17]
map[Alex:18 Kevin:15 Robert:17]
*/
