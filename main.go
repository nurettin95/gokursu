package main

import (
	"golesson/project"
)

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//conditionals.Demo3()
	//loops.Demo5()
	//arrays.Demo4()
	//functions.SelamVer("Nurettin")
	//var sonuc = functions.Topla(3, 5)
	//fmt.Println(sonuc * 10)

	//sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)

	//fmt.Println("Toplam: ", sonuc1)
	//fmt.Println("Çikarim: ", sonuc2)
	//fmt.Println("Çarpim: ", sonuc3)
	//fmt.Println("Bölüm: ", sonuc4)

	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayilar...))

	//maps.Demo1()
	//for_range.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Maindeki sayi: ", sayi)

	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki sayi: ", sayilar[0])

	//structs.Demo1()
	//structs.Demo2()

	// go goroutines.CiftSayilar()
	// go goroutines.TekSayilar()
	// time.Sleep(5 * time.Second)
	// fmt.Println("Main bitti")

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Carpim : ", carpim)

	//interfaces.Demo1()
	//interfaces.Demo2()

	//deferstatement.B()
	//deferstatement.Test()
	//deferstatement.Demo3()

	//error_handling.Demo1()

	//interfaces.Demo3()

	//error_handling.Demo2()
	//fmt.Println(error_handling.TahminEt2(102))

	//string_functions.Demo1()
	//string_functions.Demo2()

	//restful.Demo2()

	project.GetAllProducts()
	//project.AddProduct()
	//project.GetAllProducts()

}
