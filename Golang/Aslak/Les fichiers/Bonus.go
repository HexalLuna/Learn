// On se retrouve avec beaucoup de lignes de code répétables, il serait temps d'utiliser les super pouvoirs des fonctions pour mieux structurer notre code :

package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func write(text string, file *os.File) {
    if _, err := file.WriteString(text); err != nil {
        panic(err)
    }
}

func read(filename string) string {
    data, err := ioutil.ReadFile(filename)
    check(err)
    return string(data)
}

func main() {
    file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
    defer file.Close()
    check(err)

    write("Test\n", file)

    data := read(file.Name())
    fmt.Print(data)
}

/*
Résultat :

je suis un fichier de test
test
i love test
Test
*/
