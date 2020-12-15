// Pour créer une goroutine il faut placer le mot clé go avant un appel de fonction, exemple :

package main

import (
    "fmt"
    "time"
)

func run(name string) {
    for i := 0; i < 3; i++ {
        time.Sleep(1 * time.Second)
        fmt.Println(name, " : ", i)
    }
}

func main() {
    debut := time.Now()
    go run("Clément")
    go run("Robert")
    go run("Alex")
    fin := time.Now()
    fmt.Println(fin.Sub(debut))

}

/*
----------------------------------
Résultat :

Hatim  :  0
Hatim  :  1
Hatim  :  2
Robert  :  0
Robert  :  1
Robert  :  2
Alex  :  0
Alex  :  1
Alex  :  2
9.0095154s
----------------------------------
Sans grande surprise le temps d'exécution est d'un peu après 9 secondes.
*/
