//Ici nous allons voir comment vérifier le type d'une fonction
package main 

import "fmt"

type void func()

func ain()  {
	var f void

	process(func() {})
	fmt.Printf("\n %T", f)
}

func process(f void) {
	f()
}