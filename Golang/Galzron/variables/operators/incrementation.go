package main

import "fmt"

func main()  {
    var A int = 6

    A++ //Ceci est une incrémentation, je rajoute 1 à la valeur de ma variable
    fmt.Println("Vous venez d'incrémenter de 1 la valeur de la variable A, sa valeur est désormais de: ", A)

    A-- //Ceci est une décrémentation, je soustrais 1 à la valeur de ma variable
    fmt.Println("Vous venez de décrémenter de 1 la valeur de la variable A, sa valeur est désormais de: ", A)

    /**
     * Résultats:
     *
     * Incrémenration, A = 7
     * Décrémentation, A = 5
    **/
}
