package main 

import "fmt"

func main() {
    var name string
    var mdp int
	
    fmt.Println("Bonjour, veuillez vous identifiez par votre prénom s'il vous plaît:")
    fmt.Scanln(&name)

    fmt.Println("Mot de passe: ")
    fmt.Scanln(&mdp)

    if mdp == 24112004 && name == "Alexis" {
	fmt.Println("Bienvenue Alexis !")

    } else if mdp == 123456789 && name == "Clément" {
	fmt.Println("Bienvenue Clément !")

    } else {
	fmt.Println("Vous n'êtes pas identifiable dans notre 'db'")
    }	
}
