package main

import "fmt"

func main() {
    var notes = make(map[string]int)
	notes["Clément"] = 20
	notes["Alex"] = 18
	notes["Kevin"] = 15

	fmt.Println(notes)
}

/*
Résultat :

map[Alex:18 Clément:20 Kevin:15]
*/
