/*
Comme les Slices, il existe deux façons pour créer une Map.

Première méthode :
*/

package main

import (
    "fmt"
)

func main() {
    var notes map[string]int
}

// Deuxième méthode depuis la fonction make()

package main

import (
    "fmt"
)

func main() {
    var notes = make(map[string]int)
}

// Comme dans d'autres types de variables il est aussi possible de surcharger les valeurs par défaut de votre Map.

package main

import "fmt"

func main() {
    notes := map[string]int{"Clément": 20, "Alex": 18}
	fmt.Println(notes)
}

/*
Résultat :

map[Alex:18 Clément:20]
*/
