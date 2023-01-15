package variables

import "fmt"

func Demo1() {

	var metin string = "Merhaba"
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25
	fmt.Println(kdv3)

	fmt.Printf("veri tipi: %T\n", kdv3) //f formattan geliyor

	var durum bool

	var metin1 string = "Nurettin"
	var metin2 string = "Ahmet"

	durum = metin1 == metin2

	fmt.Println(durum)
	fmt.Println(!durum)
}
