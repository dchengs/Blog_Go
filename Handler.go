package main

import (
	"fmt"
	"html/template"
	"net/http"
	"./Utility"
)

//list of pages handler
func Index(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("index.html")
	if err != nil{ 
		Utility.ErrLogger(err)
	}
	fmt.Println(index.Execute(w, nil))
}

func CurrentProject(w http.ResponseWriter, r *http.Request) {
	currentProject, err := template.ParseFiles("currentProject.html")
	if err != nil{
		Utility.ErrLogger(err)
	}
	currentProject.Execute(w, nil)
}
func AboutMe(w http.ResponseWriter, r *http.Request) {
	aboutMe, err := template.ParseFiles("aboutMe.html")
	if err != nil{
		Utility.ErrLogger(err)
	}
	aboutMe.Execute(w, nil)
}


func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/currentProject/", CurrentProject)
	http.HandleFunc("/aboutMe/", AboutMe)
	http.ListenAndServe(":5000", nil)
	fmt.Println("trying")
}
