package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Kübra"

	fmt.Println(s.HasPrefix(isim, "En"))
	fmt.Println(s.HasSuffix(isim, "ra"))
	fmt.Println(s.Index(isim, "en"))

	harfler := []string{"k", "ü", "b", "r", "a"}

	sonuc := s.Join(harfler, ",")
	fmt.Println(s.Replace(sonuc, ",", "", -1))

}
