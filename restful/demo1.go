package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"user_id"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close() //işlem bittiği zaman kanalı kapat performans için

	bodyBytes, _ := ioutil.ReadAll(response.Body) //response ı anlamlı hale getirmek için (response http responsedur)hatayı almadık ondan _

	bodyString := string(bodyBytes) //string dönüşümü
	fmt.Println(bodyString)         //stringi pek kullanmayız

	var todo Todo
	json.Unmarshal(bodyBytes, &todo) //gelen body'ı todo'ya aktar
	fmt.Println(todo)
}

func Demo2() {
	todo := Todo{1, 2, "Alişveriş yapilacak", false}
	jsonTodo, err := json.Marshal(todo) //todo'yu json formatina çevirdik

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}
