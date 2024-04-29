package slices

import "fmt"

func Demo1() {
	names := make([]string, 3)
	names[0] = "Kübra"
	names[1] = "Engin"
	names[2] = "Kübra"
	//yeni eleman eklemek istediğimizde
	names = append(names, "Büşra")
	fmt.Println(names)
	fmt.Println(len(names))
}
