package logger

import (
	"log"
	"os"
)

//logging
func ErrLogger(errMessage error) {
	//Stat performs simple look up of the directory
	_, err := os.Stat("logs") 
	//checking if directory exists
	if os.IsNotExist(err){
		//creates logs directory
		err := os.Mkdir("logs", 0755)
		if err !=nil{
			log.Fatalln("Problem creating logs directory %s", err)
		}
	}
	//check if log.txt is opened, else create one
	logFile, err := os.OpenFile("logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Problem opening the file %s", err)
	}
	//ensure the file closes
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println(errMessage)
}
