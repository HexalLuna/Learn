package main 

import "fmt"

func main()  {
    var slice []int //ceci est un slice de int, en quelque sorte une array sans limite de valeur définit
    //var array []int{1, 2, 3} ==> autre manière de définir un slice

    for s := 0; s < 9; s++ {
    	slice = append(slice, s + 1) //ici je place des valeurs dans mon slice
    }
	
    for _, x := range slice {
    	fmt.Println(x)
    }

    fmt.Printf("%v", slice)
}
