package arrays

import "fmt"

func Demo2() {
	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "istanbul"
	cities[2] = "İzmir"
	cities[3] = "Adana"
	cities[4] = "Diyarbakır"
	//cities[5] tanımlanırsa dizi sınırının dışına çıkılmış olur.

	fmt.Println(cities)

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i])
	}
}
