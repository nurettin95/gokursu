package for_range

import "fmt"

func Demo2() {

	sayilar := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0

	for _, sayi := range sayilar {
		if sayi%2 == 1 {
			toplam += sayi
			fmt.Println(sayi)
		}
	}
	fmt.Println(toplam)
}
