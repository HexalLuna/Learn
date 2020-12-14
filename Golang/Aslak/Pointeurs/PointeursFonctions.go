/*
Il est concevable de déclarer des pointeurs en tant que paramètres dans une fonction.

Attention !

Il faut savoir qu'un changement de valeurs d'une variable actionné depuis un pointeur est irréversible
car le pointeur va directement modifier la valeur depuis votre adresse mémoire.
*/

package main

import "fmt"

func main() {
    var a int = 20

    fmt.Printf("Avant de modifier la variable de a : %d\n", a)

    rajouterCinq(&a)

    fmt.Printf("Après modification par des pointeurs de la variable de a : %d\n", a)

    affichage(&a)
}

// fonction qui prend en paramètre un pointeur
func rajouterCinq(pa *int) {
    *pa = *pa + 5 // modification de la valeur de la variable "a" depuis un pointeur
}

// fonction qui prend en paramètre un pointeur
func affichage(pa *int) {
    fmt.Println("affichage da valeur dans une fonction :", *pa)
}

/*
Résultat :

Avant de modifier la variable de a : 20
Après modification par des pointeurs de la variable de a : 25
affichage da valeur dans une fonction : 25
*/

/*
J'ai modifié la valeur de ma variable a depuis le pointeur défini en tant que paramètre sur la fonction rajouterCinq(),
comme je l'ai expliqué plus haut, les changements sont irréversibles peu importe l'endroit où je l'affiche (fonction main() ou autres).
*/
