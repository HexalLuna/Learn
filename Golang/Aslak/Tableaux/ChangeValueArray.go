package main

import "fmt"

func main() {
    var jours = [5]string{"Clément", "Robert", "Inconnu", "Ahmed", "Inconnu"}

    jours[0] = "Alex" // on remplace le premier element (ici Clément) par Alex
    fmt.Println(jours)

    jours = replaceByClement(jours)
    fmt.Println(jours)
}

/*
    J'utilise une fonction pour vous montrer qu'il est
    possible de prendre en paramètre un tableau 
    mais aussi de retourner un tableau dans une fonction 
*/
func replaceByClement(jours [5]string) [5]string {
    for i, jour := range jours {
        if jour == "Inconnu" {
            jours[i] = "Clément" // Remplacer "Inconnu" par "Clément"
        }
    }
    return jours
}

/*
Résultat :

[Alex Robert Inconnu Ahmed Inconnu]
[Alex Robert Clément Ahmed Clément]
*/
