package main

import (
	"dormitory-system/src/router"
	"log"
)

func main() {
	err := router.InitRouter().Run(":8090")
	if err != nil {
		log.Fatalln("Server Error: ", err)
	}
}
