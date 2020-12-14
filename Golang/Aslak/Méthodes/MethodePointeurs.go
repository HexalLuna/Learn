/*
Ça va nous permettre de modifier directement les valeurs des attributs de notre structure depuis une méthode. 
Pour le même exemple je vais créer une méthode Init() qui va initialiser les valeurs des attributs de ma structure Personnage.
*/

package main

import (
    "fmt"
)

type Personnage struct {
    nom        string
    vie        int
    puissance  int
    mort       bool
    inventaire [3]string
}

func main() {
    // initialisation de ma structure Personnage
    var p1 Personnage
    var p2 Personnage

    p1.Init("barbare", 200, 20, false, [3]string{"épée", "bouclier", "armure"})
    p2.Init("magicien", 100, 40, false, [3]string{"potions", "poisons", "bâton"})
    p1.affichage()
    p2.affichage()
}

/*
@description : Surcharge des valeurs par defaut

@return: void
*/
func (p *Personnage) Init(nom string, vie int, puissance int, mort bool, inventaire [3]string) {
    p.nom = nom
    p.vie = vie
    p.puissance = puissance
    p.mort = mort
    p.inventaire = inventaire
}

/*
@description : Affiche des informations sur un personnage

@return: void
*/
func (p Personnage) affichage() {
    fmt.Println("--------------------------------------------------")
    fmt.Println("Vie du personnage", p.nom, ":", p.vie)
    fmt.Println("Puissance du personnage", p.nom, ":", p.puissance)

    if p.mort {
        fmt.Println("Vie du personnage", p.nom, "est mort")
    } else {
        fmt.Println("Vie du personnage", p.nom, "est vivant")
    }

    fmt.Println("\nLe personnage", p.nom, "possède dans son inventaire :", p.vie)

    for _, item := range p.inventaire {
        fmt.Println("-", item)
    }
}

/*
Résultat :

--------------------------------------------------
Vie du personnage barbare : 200
Puissance du personnage barbare : 20
Vie du personnage barbare est vivant

Le personnage barbare possède dans son inventaire : 200
- épée
- bouclier
- armure
--------------------------------------------------
Vie du personnage magicien : 100
Puissance du personnage magicien : 40
Vie du personnage magicien est vivant

Le personnage magicien possède dans son inventaire : 100
- potions
- poisons
- bâton
*/
