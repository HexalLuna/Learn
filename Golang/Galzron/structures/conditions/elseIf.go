package main 

import "fmt"

func main() {
    var name string //déclaration d'une variable déstinée à contenir une chaîne de caractères
    var mdp int //déclaration de la variable servant au mot de passe
	
    fmt.Println("Bonjour, veuillez vous identifiez par votre prénom s'il vous plaît:")
    fmt.Scanln(&name) //scan ma variable name et ouvre une entrée utilisateur dans la console

    fmt.Println("Mot de passe: ")
    fmt.Scanln(&mdp) //scan ma variable mdp et ouvre une entrée utilisateur dans la console

    if mdp == 24112004 && name == "Alexis" { //vérification des variables
	fmt.Println("Bienvenue Alexis !") //informations ok

    } else if mdp == 123456789 && name == "Clément" { //vérification des variables
	fmt.Println("Bienvenue Clément !") //informations ok

    } else {
	fmt.Println("Vous n'êtes pas identifiable dans notre 'db'") //informations éronnées
    }	
}
