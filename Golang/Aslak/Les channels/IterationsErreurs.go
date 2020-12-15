package main

    import (
    "fmt"
    "time"
)

func main() {

    ch := make(chan string, 2) // buffer de taille 2

    go func() {
        ch <- "test" // 1 seul récepteur
    }()

    for elem := range ch {
        fmt.Println(elem)
    }

}

/*
Erreur :

test
fatal error: all goroutines are asleep - deadlock!

*/
/*
Alors, déjà on peut remarquer qu'on arrive à lire notre première valeur de notre channel, mais juste après nous avons un deadlock.
L'erreur vient du fait qu'on est en train de lire sur un channel avec un buffer de taille 2, 
sauf que ne nous n'avons rentré qu'une seule valeur dans notre channel d'où l'erreur.

Suite à Iterations2.go
*/
