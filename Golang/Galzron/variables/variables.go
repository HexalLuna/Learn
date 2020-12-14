package main

import "fmt"

func main()  {

	//Pour déclarer une variable et lui assigner une variable directement vous pouvez utiliser :=

	var s string //un string est une chaîne de caractères, exemple: "Aslak est gay"
	var i int //un int est un nombre entier, exemple: 100
	var f float64 //un float64 ou float32 est un nombre décimal, exemple: 86.32
	var p *int //le pointeur * est une variable dont le contenu est une adresse mémoire d'une autre variable, en bref le chemin d'accès direct de l'emplacement mémoire d'une variable
	var a []int //[] est un tableau, plus connu sous le nom d'array ou slice. C'est une variables capables de contenir une infinité de variables sur un type définit.

	fmt.Printf("%s \n%d \n%f \n%p \n%v", s, i, f, p, a)//Ici on écrit dans la console les valeurs de mes variables

	/*
	la fonction Printf() permet d'afficher plus d'informations qu'un Println() 
	il te suffit de rajouter un symbole spécial à l'endroit où tu veux afficher la variable

		%T : affiche le type d'une valeur
		%d : affiche un entier
		%s : affiche une chaîne de caractères
		%f : affiche un nombre décimal
		%b : affiche une représentation binaire

	Pour en savoir plus voici la doc de fmt: https://golang.org/pkg/fmt/	
	*/
}