package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	Id    int
	Name  string
	Check int
	Time  time.Time
}

var TodoBox []Todo

var JSON_FILE = "./file/db.json"
var index int

func makeTodo(name string) {

	index += 1
	data := Todo{Id: index, Name: name, Check: 0, Time: time.Now()}

	TodoBox = append(TodoBox, data)
	doc, _ := json.Marshal(TodoBox)
	err := ioutil.WriteFile(JSON_FILE, doc, os.FileMode(0644)) // db.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}

}

func updateTodo(checkNum int) {
	for num := range TodoBox {
		if TodoBox[num].Id == checkNum {
			TodoBox[num].Check = 1
			break
		}

	}

	doc, _ := json.Marshal(TodoBox)
	err := ioutil.WriteFile(JSON_FILE, doc, os.FileMode(0644)) // db.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}

}

func deleteTodo(deleteNum int) {

	for num := range TodoBox {
		if TodoBox[num].Id == deleteNum {

			pre_num := num
			pre_delete_num := pre_num - 1
			if num <= 0 {
				pre_delete_num = 0
			}
			TodoBox[num] = TodoBox[pre_delete_num]

		}
	}

	TodoBox = TodoBox[:len(TodoBox)-1]

	doc, _ := json.Marshal(TodoBox)
	err := ioutil.WriteFile(JSON_FILE, doc, os.FileMode(0644)) // db.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getTodoListHandler() {
	fmt.Printf("[home] Welcome to TODO list ================================\n")

	file, _ := os.Open(JSON_FILE)
	byteValue, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal([]byte(byteValue), &TodoBox)

	if TodoBox == nil {
		fmt.Printf("Empty Task ..\n Please add task.(create or add)\n")
		index = 0
		return
	}

	var check string
	var max_index = 0
	for num := range TodoBox {
		if TodoBox[num].Check == 1 {
			check = "v"
		} else {
			check = ""
		}
		if TodoBox[num].Id > max_index {
			max_index = TodoBox[num].Id
		}
		fmt.Printf("%d. (%s) %s %s \n", TodoBox[num].Id, check, TodoBox[num].Name, TodoBox[num].Time.String()[:19])

	}
	index = max_index

}

func createTodoHandler() {

	var name string
	fmt.Printf("Add item> ")
	fmt.Scanln(&name)
	if name == "!exit" {
		return
	}
	makeTodo(name)

}

func updateTodoHandler() {
	var checkNum string
	fmt.Printf("complate item> ")
	fmt.Scanln(&checkNum)
	if checkNum == "!exit" {
		return
	}
	sv, _ := strconv.Atoi(checkNum)
	updateTodo(sv)
}

func deleteTodoHandler() {
	fmt.Println(TodoBox)
	var deleteNum string
	fmt.Printf("delete item> ")
	fmt.Scanln(&deleteNum)

	if deleteNum == "!exit" {
		return
	}

	dnum, _ := strconv.Atoi(deleteNum)

	deleteTodo(dnum)

}

func printer(command string) {

	fmt.Printf("[%s] Welcome to TODO list ================================\n", command)

}

func handler(command string) {

	switch {
	case command == "create" || command == "add":
		printer(command)
		createTodoHandler()

	case command == "update" || command == "complate":
		printer(command)
		updateTodoHandler()

	case command == "delete":
		printer(command)
		deleteTodoHandler()

	case command == "!exit":
		getTodoListHandler()

	default:
		fmt.Println("Not support command ")
		return

	}
	fmt.Printf("===================================================\n")
}

func main() {

	fmt.Println("command : create,  update(complate), delete")

	for {
		getTodoListHandler()
		var command string
		fmt.Printf("Enter> ")
		fmt.Scanln(&command)
		if command == "!exit" {
			os.Exit(0)
		}
		handler(command)
	}

}

// [{id : 2, "name": "할일", "do": 1, "data": time },{id : 2, "name": "22", "do": 0, "data": time }]
