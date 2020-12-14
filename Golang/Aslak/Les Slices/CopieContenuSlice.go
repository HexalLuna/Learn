/*
Il est possible de copier le contenu d'une slice source vers une slice cible grâce à la fonction copy(). 
Cette fonction prend comme premier paramètre la slice cible et comme deuxième paramètre la slice source.

Attention

La fonction copy() copie le contenu d’une tranche source vers une source cible, 
il est donc important que la slice cible soit de même taille que la slice source afin de copier le contenu total de la cible.
*/

package main

import "fmt"

func main() {

	animaux1 := []string{"Lion", "Cheval", "Ours"}
    fmt.Println("Contenu du tableau animaux1 :", animaux1)

	// création d'une slice cible avec la même taille que la slice source
    animaux2 := make([]string, len(animaux1))

	// copie du contenu de la slice source vers la slice cible
	copy(animaux2, animaux1)
	
    fmt.Println("Contenu du tableau animaux2 :", animaux2)
}

/*
Résultat :

Contenu du tableau animaux1 : [Lion Cheval Ours]
Contenu du tableau animaux2 : [Lion Cheval Ours]
*/
