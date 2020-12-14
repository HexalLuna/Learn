package main

import "fmt"

type user struct {
	name string
	todos []*todo
}

type todo struct {
	text string
	done bool
}

type printable interface {
	print() string
}

func (u user) show()  {
	fmt.Printf("Name: %s\n", u.name)
	fmt.Printf("Todos: %d\n", len(u.todos))

	for _, t := range u.todos {
		fmt.Printf("<Todo text=%s done=%v", t.text, t.done)
	}
}

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
}