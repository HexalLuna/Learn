// Dans cet exemple je vais créer un channel avec 5 récepteurs et 5 expéditeurs :

package main

import (
    "fmt"
    "sync"
    "time"
)

var wg = sync.WaitGroup{}

func main() {

    now := time.Now()
    ch := make(chan int)

    // 5 expéditeurs
    for j := 0; j < 5; j++ {
        wg.Add(1)
        go func() {
            time.Sleep(time.Second * 2)
            i := <-ch
            fmt.Println(i)
            wg.Done()
        }()
    }

    // 5 récepteurs
    for j := 0; j < 5; j++ {
        wg.Add(1)
        go func() {
            time.Sleep(time.Second * 2)
            ch <- 50

            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println(time.Now().Sub(now))
}

/*
Résultat :

50
50
50
50
50
2.000827906s
*/
