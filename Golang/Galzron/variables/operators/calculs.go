package main

import "fmt"

func main()  {

    var a int = 23
    var b int = 5

    c := a + b //les valeurs de a et b s'additionne dans une nouvelle variable appelée c
    a += b //La valeur de b s'assigne a celle de la variable a
    fmt.Println("a + b = ", c, "and a += b = ", a)

    //Nous allons maintenant réitérer cette opération mais avec les autres signes opératoires existants et en lister d'autres
    d := a - b
    a -= b
    fmt.Println("a - b = ", d, "and a -= b = ", a)

    e := a * b
    a *= b
    fmt.Println("a * b = ", e, "and a *= b = ", a)

    f := a / b
    a /= b
    fmt.Println("a / b = ", f, "and a /= b = ", a)

    g := a % b
    a %= b
    fmt.Println("a % b = ", g, "and a %= b = ", a)
}

/**
 *  * == multiplication
 *  / == division
 *  - == soustraction
 *  + == addition
 *  % == modulo ( reste de la division)
 *  << == décalage du bit vers la gauche
 *  >> >>> == déclagae du bit vers la droite
 *  & == ET binaire
 *  | == OU binaire
 *  ^ == OU EXCLUSIF binaire
**/

/**
 * Operation results:
 *
 * a + b =  28 and a += b =  28
 * a - b =  23 and a -= b =  23
 * a * b =  115 and a *= b =  115
 * a / b =  23 and a /= b =  23
 * a % b =  3 and a %= b =  3
**/
