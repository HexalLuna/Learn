package main

import (
	"fmt"
)

func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}

	for i := 0; i < len(jours); i++ {
		fmt.Println(jours[i], "est le jour numéro", (i + 1), "de la semaine")
	}
}

/*
Résultat :

lundi est le jour numéro 1 de la semaine
mardi est le jour numéro 2 de la semaine
mercredi est le jour numéro 3 de la semaine
jeudi est le jour numéro 4 de la semaine
vendredi est le jour numéro 5 de la semaine
samedi est le jour numéro 6 de la semaine
dimanche est le jour numéro 7 de la semaine
*/
