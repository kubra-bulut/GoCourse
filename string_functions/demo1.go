package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Engin"
	fmt.Println(s.Count(isim, "E"))
	fmt.Println(s.Contains(isim, "a"))
	sonuc := s.Index(isim, "i")

	if sonuc != -1 {
		fmt.Println("i harfi var")
	} else {
		fmt.Println("i harfi yok")
	}
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
