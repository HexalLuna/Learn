package main

import (
    "fmt"
    "sync"
    "time"
)

var wg = sync.WaitGroup{}

func main() {

    now := time.Now()
    const size int = 10
    ch := make(chan int, size) // channel avec un buffer de taille 10

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

    // 10 récepteurs
    for j := 0; j < size; j++ {
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
2.00064941s
*/
