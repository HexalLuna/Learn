package main

import "fmt"

func main() {
    var a int = 4
    var b int = 2

	a += b
	fmt.Println("a += b  = ", a)

	a -= b
	fmt.Println("a -= b  = ", a)

	a *= b
	fmt.Println("a *= b  = ", a)

	a /= b
	fmt.Println("a /= b  = ", a)

	a %= 3
	fmt.Println("a %= b  = ", a)
}

/* Résultat :

a += b  =  6
a -= b  =  4
a *= b  =  8
a /= b  =  4
a %= b  =  1
*/
