package string_functions

//alias
import (
	"fmt"
	s "strings"
)

// case sensitive -Büyük küçük harf duyarlı
// ascii
func Demo1() {
	isim := "Nurettin"
	fmt.Println(s.Count(isim, "n"))
	fmt.Println(s.Contains(isim, "n"))
	fmt.Println(s.Index(isim, "n")) //bulamazsa -1 döndürür
	sonuc := s.Index(isim, "a")

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
