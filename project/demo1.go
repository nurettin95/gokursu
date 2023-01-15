package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) //response.Body oku diyip bodyBytes'a atadık

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	/*	bodyBytes'ı bellek adresi ile beraber produtsa atıyoruz
		yani products bu bağlamda  'var produts'ın' aynı adresi kullanmış olacak dolayısıyla bu da değişmiş olacak.
		Aksi taktirde & işareti yollamasak aslında yepyeni bir product nesnesi gibi gidiyor 'var products' değişmeyecek olacaktı
	*/
	fmt.Println(products)

}

func AddProduct() {
	//product := Product{Id: 4, ProductName: "Telephone", CategoryId: 1, UnitPrice: 6000.0}
	product := Product{ProductName: "Microphone", CategoryId: 1, UnitPrice: 2000.0}

	//bu bir struct json api struct'ı tanımaz o yüzden json'a çevirmemiz lazım Marshal kütüphanesi ile
	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) //data dönecek

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse) //bodyBytes ı productResponse'a ata
	fmt.Println("Kaydedildi: ", productResponse)
}
