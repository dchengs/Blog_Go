package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"os"
)

//list of pages handler
func index(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("index.html")
	if err != nil{
		errLogger(err)
	}
	fmt.Println(index.Execute(w, nil))
}

func currentProject(w http.ResponseWriter, r *http.Request) {
	currentProject, err := template.ParseFiles("currentProject.html")
	if err != nil{
		errLogger(err)
	}
	currentProject.Execute(w, nil)
}
func aboutMe(w http.ResponseWriter, r *http.Request) {
	aboutMe, err := template.ParseFiles("aboutMe.html")
	if err != nil{
		errLogger(err)
	}
	aboutMe.Execute(w, nil)
}

//logging
func errLogger(errMessage error){
	//make log directory if it doesn't exist
	err := os.Mkdir("logs",0755)
	if err != nil{
		log.Fatalln("Problem creating path %v", err)
	}
	//check if log.txt is opened, else create one
	logFile,err := os.OpenFile("logs/log.txt", os.O_RDWR | os.O_CREATE |os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln("Problem opening the file %v", err)
	}
	//ensure the file closes
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println(errMessage)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/currentProject/", currentProject)
	http.HandleFunc("/aboutMe/", aboutMe)
	http.ListenAndServe(":5000", nil)
	fmt.Println("trying")
}
