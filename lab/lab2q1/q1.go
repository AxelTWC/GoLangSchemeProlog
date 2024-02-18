package main

import "fmt"

type dog struct {
	name   string
	race   string
	female bool
}

func (d *dog) rename(name string) {
	d.name = name
}
func main() {
	fido := dog{"Fido", "Poodle", false}
	fido.rename("Cocotte")
	fmt.Print(fido)
}
