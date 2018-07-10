package main

import (
	"fmt"
	"html/template"
	"net/http"
	"crypto/sha256"
	"./Utility"
)

//list of pages handler
func Index(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("html/index.html")
	if err != nil {
		Utility.ErrLogger(err)
	}
	fmt.Println(index.Execute(w, nil))
}

func CurrentProject(w http.ResponseWriter, r *http.Request) {
	currentProject, err := template.ParseFiles("html/currentProject.html")
	if err != nil {
		Utility.ErrLogger(err)
	}
	currentProject.Execute(w, nil)
}
func AboutMe(w http.ResponseWriter, r *http.Request) {
	aboutMe, err := template.ParseFiles("html/aboutMe.html")
	if err != nil {
		Utility.ErrLogger(err)
	}
	aboutMe.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request){
	if(r.Method() == "GET"){
		loginPage, err = := template.ParseFiles("html/login.html")
		if err != nil{
			Utility.ErrLogger(err)
		}
	}
	else{
		username:=r.Form["username"]
		password:= r.Form["password"]
		//encrypt password input
		sha_256 := sha256.New()
		//convert password to byte array
		sha_256.Write([]byte(password))
		encryptedPW := hex.EncodeToString(sha_256.Sum(nil))
		//check with server

	}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/currentProject/", CurrentProject)
	http.HandleFunc("/aboutMe/", AboutMe)
	http.ListenAndServe(":5000", nil)
	fmt.Println("trying")
}
