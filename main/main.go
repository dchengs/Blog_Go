package main

import (
	"Blog_go/utilities/logger"
	"fmt"
	"html/template"
	"net/http"
)

//list of pages handler
func Index(w http.ResponseWriter, r *http.Request) {
	//shows 404 not found
	if r.URL.Path != "/" {
		NotFound(w, r)
	} else {
		template, err := template.ParseFiles("../htmls/index.html")
		if err != nil {
			logger.ErrLogger(err)
		} else {
			template.Execute(w, nil)
		}
	}
}

func Project(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../htmls/currentProject.html")
	if err != nil {
		logger.ErrLogger(err)
	} else {
		template.Execute(w, nil)
	}
}
func AboutMe(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../htmls/aboutMe.html")
	if err != nil {
		logger.ErrLogger(err)
	} else {
		template.Execute(w, nil)
	}
}

func Resume(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../htmls/resume.html")
	if err != nil {
		logger.ErrLogger(err)
	} else {
		template.Execute(w, nil)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	//user requests the login page

	if r.Method == "GET" {
		template, err := template.ParseFiles("../htmls/login.html")
		if err != nil {
			logger.ErrLogger(err)
		}
		template.Execute(w, nil)
	} else {
		//simple debug testing
		debugUsername := "chengs"
		debugPassword := "123123123"
		r.ParseForm()
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		if debugUsername == username && debugPassword == password {
			fmt.Println("Correct password")
		} else {
			fmt.Println("Incorrect password")
		}
		/*
			//encrypt password input
			sha_256 := sha256.New()
			//convert password to byte array
			sha_256.Write([]byte(password))
			encryptedPW := hex.EncodeToString(sha_256.Sum(nil))
			//check with serve
		*/

	}
}
func NotFound(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../htmls/404.html")
	if err != nil {
		logger.ErrLogger(err)
	} else {
		template.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/project/", Project)
	http.HandleFunc("/aboutMe/", AboutMe)
	http.HandleFunc("/login/", Login)
	http.HandleFunc("/resume/", Resume)
	http.ListenAndServe(":5000", nil)
}
