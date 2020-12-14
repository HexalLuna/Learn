package main

import "fmt"

func main() {
    var a int = 20 //  je déclare une variable nommée a de type int
    var ap *int // création d'un pointeur | je déclare juste un pointeur nommé ap de type int
    ap = &a // stockage de l'adresse mémoire de la variable dans notre pointeur
    /*
    Petit rappel "un pointeur ne peut stocker que des adresses mémoire".
    Ça tombe bien car ici grâce au signe & on peut accéder à l'adresse mémoire de notre variable a
    */

    fmt.Printf("Adresse de votre variable a : %p\n", &a) // Ici j’affiche tout simplement l’adresse mémoire de la variable a

    fmt.Printf("Valeur de votre variable (pointeur) ap : %p\n", ap) // j'affiche la valeur du pointeur ap (il possède bien comme valeur l’adresse de votre variable a).

    fmt.Printf("Valeur de l'adresse %p: %d\n", ap, *ap) // j’accède directement à la valeur de la variable a depuis son adresse grâce au pointeur ap.
    /*
    L'astérisque * nous permet entre autres de déclarer un pointeur mais il permet aussi d'accéder à la valeur de la variable pointée, 
    sur cette ligne de code on l'utilise pour accéder à la valeur contenue dans l'adresse mémoire de la variable a
    */
}

/*
Résultat :

Adresse de votre variable: 0xc000056058
Valeur de votre variable (pointeur) ap : 0xc000056058
Valeur de l'adresse 0xc000056058: 20
*/
