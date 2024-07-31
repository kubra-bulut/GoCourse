package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	//variables.Demo1()
	//conditionals.Demo3()
	//loops.Demo6()
	//slices.Demo2()

	//var sonuc = functions.Topla(5, 6)
	//fmt.Println(sonuc)

	//sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	//fmt.Println("toplam: ", sonuc1)
	//fmt.Println("Çıkarım: ", sonuc2)
	//fmt.Println("Çarpım: ", sonuc3)
	//fmt.Println("Bölüm: ", sonuc4)

	fmt.Println(functions.ToplaVariadic(5, 3, 2, 1))
}
