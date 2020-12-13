// - Récupérer tous les éléments de notre tableau :

package main

import (
	"fmt"
)

func main() {
	var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
	fmt.Println(jours[:]) // on récupère tous les éléments
}
/*
Résultat :

[lundi mardi mercredi jeudi vendredi samedi dimanche]
*/

// - Récupérer les trois premiers éléments de notre tableau :

package main

import (
    "fmt"
)

func main() {
    var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
    fmt.Println(jours[:3]) // trois premiers éléments
}

/*
Résultat :

[lundi mardi mercredi]
*/

// - Récupérer tous les résultats à partir du troisième index :

package main

import (
    "fmt"
)

func main() {
    var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
    fmt.Println(jours[3:])
}
/*
Résultat :

[jeudi vendredi samedi dimanche]
*/

// - Récupérer les résultats du premier jusqu'au troisième index (exclus) :

package main

import (
    "fmt"
)

func main() {
    var jours = [7]string{"lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi", "dimanche"}
    fmt.Println(jours[1:3])
}
/*
Résultat :

[mardi mercredi]
*/
