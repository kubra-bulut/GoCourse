package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) Save() {
	fmt.Println("eklendi", c.firstName)
}

func (c customer) Update() {
	fmt.Println("güncellendi", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Engin", lastName: "Demiroğ", age: 35}
	c.Save()
	c.Update()
}
