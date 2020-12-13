package main

import (
    "fmt"
)

func main() {
    var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

    fmt.Println("jours[0] =", jours[0])
    fmt.Println("jours[1] =", jours[1])
    fmt.Println("jours[2] =", jours[2])
    fmt.Println("jours[3] =", jours[3])
    fmt.Println("jours[4] =", jours[4])
    fmt.Println("jours[5] =", jours[5])
    fmt.Println("jours[6] =", jours[6])
}

/*
Résultat :

jours[0] = lundi
jours[1] = mardi
jours[2] = mercredi
jours[3] = jeudi
jours[4] = vendredi
jours[5] = samedi
jours[6] = dimanche
*/
