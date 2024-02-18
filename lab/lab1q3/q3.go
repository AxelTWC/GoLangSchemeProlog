package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Person struct {
	firstName string
	lastName  string
	iD        int
}

func inPerson(p *Person, id int) (nextId int, err error) {
	fmt.Print("Input the First Name: ")
	reader := bufio.NewReader(os.Stdin)
	firstName, err := reader.ReadString('\n')

	if len(firstName) == 2 {
		return nextId, errors.New("WRONG MESSAGE")
	}
	fmt.Print("Input the Last Name: ")
	lastName, err := reader.ReadString('\n')

	if len(lastName) == 2 {
		return nextId, errors.New("WRONG MESSAGE")
	}
	*p = Person{firstName, lastName, id}
	nextId = id
	p.iD = nextId
	nextId += 1
	return
}
func printPerson(p Person) {
	fmt.Println("_______________________________")
	fmt.Print("First Name: ", p.firstName)
	fmt.Print("Last Name: ", p.lastName)
	fmt.Println("ID: ", p.iD)
	fmt.Println("_______________________________")
}
func main() {
	nextId := 101
	for {
		var (
			p   Person
			err error
		)
		nextId, err = inPerson(&p, nextId)
		if err != nil {
			fmt.Println("Invalid entry ... exiting")
			break
		}
		printPerson(p)
	}
}
