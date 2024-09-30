package defer_statement

import "fmt"

func A() {
	fmt.Println("A foknsiyonu çalıştı")
}
func B() {
	//en son bu fonk çalışacak
	//lifo mantığı ile çalışır
	defer A()
	defer C()
	fmt.Println("B foknsiyonu çalıştı")

}

func C() {
	fmt.Println("C foknsiyonu çalıştı")
}
