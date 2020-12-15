// Voir le README.md pour en savoir plus.

package main

import (
	"fmt"
	"time"
)

func run(ch chan string, name string) {
	time.Sleep(time.Second * 2)
	fmt.Println("fonction run() :", name)
	ch <- name
}

func main() {

	now := time.Now()

	ch := make(chan string)

	go run(ch, "channel 1")
	fmt.Println("fonction main() :", <-ch)

	go run(ch, "channel 2")
	fmt.Println("fonction main() :", <-ch)

	fmt.Println(time.Now().Sub(now))
}

/*
Résultat :

fonction run() : channel 1
fonction main() : channel 1
fonction run() : channel 2
fonction main() : channel 2
4.00097503s

/*
Nous avons lancé 2 fois la goroutine de la fonction run() avec une lecture sur le channel ch,
nous pouvons remarquer que le temps d’exécution est de 4 secondes avec un intervalle de 2 secondes entre chaque goroutine, 
ce qui prouve que les channels sont bien bloquants.
*/
