package main 

import "fmt"

type user struct { //Définition du mon type user qui est une structure de donnée
   id int 
   name string
}

type entier int //Création d'un nouveau type
type Entier = int //Création d'un aliase vers un type

func main()  {
    u := user{id: 18, name: "Clément"} //attribution des valeurs sur mon type user
    u2 := user{id: 16, name: "Alexis"}

    var x int = 10
    var z Entier = 12

    var res entier
    res = somme(x, z)

    fmt.Printf("%v \n%v", u, u2)
    fmt.Printf("%T\n %T\n %T\n %v\n", x, z, res, res)
}

//func () identifier(params) (returns)
func somme(a, b int) entier {
    return entier(a + b)
}
