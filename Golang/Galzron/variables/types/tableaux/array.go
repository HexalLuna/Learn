package main 

import "fmt"

func main()  {
	var array [5]int //mon tableau va prendre 5 éléments

	for i := 0; i < 5; i++ {
		array[i] = i + 1 //ici je place des valeurs dans mon array
	}

	fmt.Printf("%v", array)
}