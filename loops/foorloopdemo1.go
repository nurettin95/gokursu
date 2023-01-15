package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dunya!"

	i := 1

	for i <= 5 {
		fmt.Println(metin)
		i++
	}

	fmt.Println("Bitti.")
}
