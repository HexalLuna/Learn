/*
Pour déclarer une interface nous avons besoin du mot-clé type qui permet d'associer un nom à votre interface et du mo-clé interface 
qui indique à votre compilateur qu'il s'agit d'une interface.
*/
package main

import (
    "fmt"
)

type Forme interface { // création de L'interface Forme
    Air() float64 // signature de la méthode Air()
    Perimetre() float64 // signature de la méthode Perimetre()
}

func main() {
    var f Forme // Initialisation de l'interface Forme
    fmt.Println(f)
}

*/
Comme vous pouvez le voir dans une interface on ne met que des signatures de méthodes sans aucun attributs. 
Pour le moment ces méthodes ne sont pas encore utilisées car pour l'instant je vous montre juste comment déclarer une interface.
*/
/*
-----------------------------
Résultat :

nil
-----------------------------
D'après le résultat ci-dessus, nous pouvons voir que la valeur de notre interface est nil 
car il n'existe aucune structure qui implémente l'interface Forme.
*/
