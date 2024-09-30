package main

import (
	"golesson/interfaces"
)

func main() {
	//variables.Demo1()
	//conditionals.Demo3()
	//loops.Demo6()
	//slices.Demo2()

	//----FUNCTIONS-------
	//var sonuc = functions.Topla(5, 6)
	//fmt.Println(sonuc)

	//sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	//fmt.Println("toplam: ", sonuc1)
	//fmt.Println("Çıkarım: ", sonuc2)
	//fmt.Println("Çarpım: ", sonuc3)
	//fmt.Println("Bölüm: ", sonuc4)

	// fmt.Println(functions.ToplaVariadic(5, 3, 2, 1))
	//maps.Demo1()

	//----POINTERS-------
	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Main:", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Main:", sayilar[0])

	//----STRUCT-------
	//structs.Demo2()

	//Routine and Channels
	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım: ", carpim)

	//Interfaces
	//interfaces.Demo1()
	//interfaces.Demo2()

	//Defer
	//defer_statement.Test()
	//defer_statement.Demo3()
	//error_handling.Demo1()

	interfaces.Demo3()
}
