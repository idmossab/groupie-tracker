package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	
	byteValue,err :=os.ReadFile("users.json")
	if err!=nil{
		log.Fatal("Error reading Json file :",err)
	}
	var users []User
	//fech data json to struct
	json.Unmarshal(byteValue,&users)

	temp :=template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/",func (w http.ResponseWriter,r *http.Request){
		temp.Execute(w,users)
	})

	http.ListenAndServe(":8080",nil)
}
