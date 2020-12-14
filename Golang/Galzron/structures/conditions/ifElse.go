package main 

import "fmt"

func main()  {
	var name string = "Alexis"
	var age int = 16

	if name == "Alexis" && age == 16 { //si name est égale à Alexis et age a 16
		fmt.Println("Bienvenue Alexis !") //J'affiche un message de bienvenue
	} else {
		fmt.Println("Vous n'êtes pas Alexis") //sinon j'affiche un autre message
	}
}