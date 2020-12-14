package main

import "fmt"

func main() {
    notes := map[string]int{"Clément": 20, "Alex": 18}

	fmt.Println("La note de Clément est :", notes["Clément"])
	fmt.Println("La note de Alex est :", notes["Alex"])
}

/*
Résultat :

La note de Clément est : 20
La note de Alex est : 18
*/
