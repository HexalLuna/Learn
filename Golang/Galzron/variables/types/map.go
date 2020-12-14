package main

import "fmt"

func main3()  {
	h := make(map[string]int)
	h["Alexis"] = 16
	
	res, found := h["Alexis"]

	fmt.Println(res, found)
	fmt.Printf("%v \n %d", h, h["Alex"]) //Ici j'affiche le type de ma variable h
}