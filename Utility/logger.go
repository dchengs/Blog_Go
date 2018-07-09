package Utility

import (
	"log"
	"os"
)

//logging
func ErrLogger(errMessage error) {
	//make log directory if it doesn't exist
	err := os.Mkdir("logs", 0755)
	if err != nil {
		log.Fatalln("Problem creating path %v", err)
	}
	//check if log.txt is opened, else create one
	logFile, err := os.OpenFile("logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Problem opening the file %v", err)
	}
	//ensure the file closes
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println(errMessage)
}
