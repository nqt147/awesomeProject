package main

import (
	s "github.com/awesomeProject/service"
	"log"
)

func main() {
	log.Print("Run success!")
	router := &s.Router{}
	router.InitializeAPIConfig()
	router.Run(":8080")
}
