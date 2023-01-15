package for_range

import "fmt"

func Demo1() {

	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}

	for i, sehir := range sehirler {
		fmt.Print(i, " ")
		fmt.Println(sehir)
	}

}
