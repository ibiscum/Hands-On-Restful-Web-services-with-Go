package main

import (
	"log"

	"github.com/ibiscum/Hands-On-Restful-Web-services-with-Go/chapter07/basicExample/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database tables are successfully initialized.")
}
