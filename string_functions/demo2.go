package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Nurettin"
	fmt.Println(s.HasPrefix(isim, "Nur")) // Nur ile başlıyor mu
	fmt.Println(s.HasSuffix(isim, "in"))  // in ile bitiyor mu
	fmt.Println(s.Index(isim, "ti"))
	harfler := []string{"n", "u", "r", "e", "t", "t", "i", "n"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", -1)) // -1 tümünü değiştir
	fmt.Println(s.Split(sonuc, "*"))
	fmt.Println(s.Repeat(sonuc, 5))
}
