package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo2.txt")
	//dosya bulunursa error=nil

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı: ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamadı")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
