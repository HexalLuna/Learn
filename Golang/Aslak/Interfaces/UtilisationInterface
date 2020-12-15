/*
Notre interface Forme possède deux signatures :
- Air() qui retourne un type float64
- Perimetre() qui retourne un type float64
La structure qui va implémenter l'interface Forme doit obligatoirement implémenter les méthodes Air() et Perimetre() signées dans l'interface Forme.
*/

package main

import (
    "fmt"
)

type Forme interface {
    Air() float64
    Perimetre() float64
}

type Rectangle struct {
    largeur  float64
    longueur float64
}

/* 
Pour implémenter une interface dans Go, il suffit
d'implémente toutes les méthodes de l'interface. Ici on
implémente la méthode Air() de l'interface Forme.
 */
func (r Rectangle) Air() float64 {
	return r.largeur * r.longueur
}

/*
On implémente la méthode Perimetre() de l'interface Forme
*/
func (r Rectangle) Perimetre() float64 {
	return 2 * (r.largeur * r.longueur)
}


func main() {
    var f Forme
    f = Rectangle{5.0, 4.0} // affectation de la structure Rectangle à l'interface Forme 
    r := Rectangle{5.0, 4.0} 
    fmt.Println("Type de f :", f)
    fmt.Printf("Valeur de f : %v\n", f)
    fmt.Println("Air du rectangle r :", f.Air())
    fmt.Println("f == r ? ", f == r)
}

// Il est tout à fait possible d'affecter à notre interface Forme la structure Rectangle puisque la structure Rectangle implémente l'interface Forme.

/*
Résultat :

Type de f : {5 4}
Valeur de f : {5 4}
Air du rectangle r : 20
f == r ?  true
*/
// Sur le résultat on peut remarquer que f == r ? est à true car dans mon exemple f et r ont le même type et la même valeur.
