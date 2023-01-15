package deferstatement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalisti")
}

func C() {
	fmt.Println("c fonksiyonu çalisti")
}

func D() {
	fmt.Println("d fonksiyonu çalisti")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("b fonksiyonu çalisti")
}

//LIFO son giren ilk çıkar
