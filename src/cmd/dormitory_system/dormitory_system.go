package main

import (
	"dormitory-system/src/database"
	"dormitory-system/src/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading env file: ", err)
	}
	database.ConnectRedis()
	database.ConnectMysql()
	err = router.InitRouter().Run(":8090")
	if err != nil {
		log.Fatalln("Server Error: ", err)
	}
}
