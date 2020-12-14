//Ici nous allonrs voir les bases d'une structures ainsi que les interfaces
package main

import "fmt"

type user struct { //type
    name string
    todos []*todo
}

type todo struct {
    text string
    done bool
}
type todolist []*todo //todolist est un slice de pointeur todo

func(tl todolist) print() string {
    var str string
    str = fmt.Sprintf("Type : TodoList\nTaille : %d\n", len(tl))
    str = fmt.Sprintf("%s\n%s", str, "-----------------------------------------------")
    for idx, t := range tl {
    	str = fmt.Sprintf("%s\n[%d]\n%s", str, idx+1, t.text)
    }
    return str
}

type printable interface { //Tout type qui a une fonction print() qui return string est un type qui satissfait l'interface printable
    print() string
}

func (u user) show()  {
    fmt.Printf("Name: %s\n", u.name)
    fmt.Printf("Todos: %d\n", len(u.todos))

    for _, t := range u.todos {
    	fmt.Printf("<Todo text=%s done=%v", t.text, t.done)
    }
}

//func() with receiver
//func (receiver) identifier(args) (returns)
func (u *user) addTodo(t *todo)  {
    u.todos= append(u.todos, t)
}

func (t *todo) toggle() {
    t.done = !t.done
}

func (u user) print() string {
    return fmt.Sprintf("Hello User, My name is %s.\n", u.name)
}

func details(t printable)  {
    fmt.Println(t.print())
}

func (t todo) print() string {
    return fmt.Sprintf("<Todo text=%s done=%v>\n", t.text, t.done)
}

func main()  {

    u := user{name: "Alexis"}
    t := todo{text: "Finir le bot Miraï"}
	
    t.toggle()
    u.addTodo(&t)

    u.show()
    t.toggle()
    u.show()

    details(u)
    details(t)

    tl := todolist(u.todos)
    details(tl)
}

//Les interfaces sont des moyens super puissant pour supporter le polymorphisme, définir un comportement que l'on peut vouloir généralisé a différent type
//On peut définir les fonctions qui satisfont les conditions d'une interfaces pour une strcture que nous abons défini nous mêmes, ou un type que l'on a défini voir déjà existant dans GO
