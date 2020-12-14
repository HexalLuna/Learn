/*
Pour rajouter un élément dans votre slice il faut utiliser la fonction append(), 
qui prend comme paramètres d'abord votre slice et ensuite l'élément que vous voulez rajouter et elle vous retournera une nouvelle Slice avec l'élément rajouté.
*/

package main

import "fmt"

func main() {
    var mois []string

    mois = append(mois, "Janvier")
    fmt.Println(mois)

    mois = append(mois, "Février")
    fmt.Println(mois)
}

/*
Résultat :

[Janvier]
[Janvier Février]
*/
