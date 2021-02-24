package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

	
type Todo struct {
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

type Todos []Todo


func jSONToTab(pathToJSON string) Todos{
	file, err := ioutil.ReadFile(pathToJSON)
    if err != nil {
      fmt.Print(err)
	}
	
	sb := string(file)

	data := Todos{}
	_ = json.Unmarshal([]byte(sb), &data)

	return data

}

func listTodo(w http.ResponseWriter, r *http.Request){
	data := jSONToTab("data/todo.json")
	
	json.NewEncoder(w).Encode(data)
}

func addTodo(w http.ResponseWriter, r *http.Request){
	action := r.FormValue("action")
	data := jSONToTab("data/todo.json")
	newStruct := &Todo{
		Todo: action,
	}

	data = append(data, *newStruct)

	dataBytes, err := json.MarshalIndent(data, "", "    ")


	err = ioutil.WriteFile("data/todo.json", dataBytes, 0644)
	if err != nil {
        fmt.Println("eroor")
	}
	
	json.NewEncoder(w).Encode("{sucess: true}")
}

func checkTodo(w http.ResponseWriter, r *http.Request){
	action := r.FormValue("action")
	data := jSONToTab("data/todo.json");

	var index int

	for i:= 0; i < len(data); i++{
		if(action == data[i].Todo){
			index = i
			
		}
	}

	data[index].Done = true

	dataBytes, err := json.MarshalIndent(data, "", "    ")

	err = ioutil.WriteFile("data/todo.json", dataBytes, 0644)
	if err != nil {
        fmt.Println("error")
    }
}

func deleteTodo(w http.ResponseWriter, r *http.Request){
	action := r.FormValue("action")
	data := jSONToTab("data/todo.json")


	for i:= 0; i < len(data); i++{
		if(action == data[i].Todo){
			data= append(data[:i], data[i+1:len(data)]...)

		}
	}
	
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	err = ioutil.WriteFile("data/todo.json", dataBytes, 0644)
	if err != nil {
        fmt.Println("error")
    }

}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("endpoint reached")
}

func handleRequest(){
	fmt.Println("sever running at http://localhost:8081/")

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/list", listTodo).Methods("GET")
	myRouter.HandleFunc("/add", addTodo).Methods("POST")
	myRouter.HandleFunc("/check", checkTodo).Methods("POST")
	myRouter.HandleFunc("/sup", deleteTodo).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}


func main()  {
	handleRequest()

}