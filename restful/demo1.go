package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ToDo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodybytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodybytes)
	fmt.Println(bodyString)

	var todo ToDo
	json.Unmarshal(bodybytes, &todo)
	fmt.Println(todo)
}

func Demo2() {
	todo := ToDo{1, 2, "Alışveriş Yapılacak", false}
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=UTF-8", bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	bodybytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodybytes)
	fmt.Println(bodyString)

	var todoResponse ToDo
	json.Unmarshal(bodybytes, &todoResponse)
	fmt.Println(todoResponse)
}
