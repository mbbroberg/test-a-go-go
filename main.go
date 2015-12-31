package main

import "fmt"

type Person struct {
	Name string
	Age int8
}

func (p *Person) Speak() {
	fmt.Println("Hey there.")
}

func birthPerson(n string) Person {
	//Person.Name := n
	//Person.Age = 0
	fmt.Printf("You made a person named %s\n", n)
	return Person{Name: n, Age: 0}
}

func main() {
	fmt.Println("Running this program")
	gary := birthPerson("Gary")
	fmt.Printf("%s says: ", gary.Name)
	gary.Speak()
}
