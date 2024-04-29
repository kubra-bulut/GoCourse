package slices

import "fmt"

func Demo2() {
	cities := []string{"Ankara", "İstanbul", "Düzce"}
	fmt.Println(cities)
	citiesCopy := make([]string, len(cities))

	copy(citiesCopy, cities)
	fmt.Println(citiesCopy)

	cities = append(cities, "Adana")
	fmt.Println(cities)
	fmt.Println(citiesCopy)
	//kopyalama yapıldıktan sonra 2 slice birbirinden farklı davranır.

	fmt.Println(cities[1:3])
}
