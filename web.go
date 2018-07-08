package main 

import (
	"fmt"
	"net/http"
	"html/template"
)

//list of pages handler
func index(w http.ResponseWriter, r *http.Request){
	index, err := template.ParseFiles("index.html")
	fmt.Println(err)
	fmt.Println(index.Execute(w, nil))
}

func currentProject(w http.ResponseWriter, r *http.Request){
	currentProject, err = template.ParseFiles("currentProject.html")
	fmt.Println(err)
	currentProject.Execute(w,nil)
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/currentProject/", currentProject)
	http.ListenAndServe(":5000", nil)
	fmt.Println("trying")
}