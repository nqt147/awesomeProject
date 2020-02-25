package main

import (
	s "../service"
	"log"
)

func main() {
	log.Print("Run success!")
	router := &s.Router{}
	router.InitializeAPIConfig()
	router.Run(":8080")
	log.Print("end")
}
