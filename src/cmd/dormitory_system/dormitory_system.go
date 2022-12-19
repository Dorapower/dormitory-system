package main

import (
	"dormitory-system/src/database"
	"dormitory-system/src/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error when loading env file: ", err)
	}
	database.ConnectRedis()
	database.ConnectMysql()
	// rabbitmq.Connect()
	port := os.Getenv("API_PORT")
	err = router.InitRouter().Run(":" + port)
	if err != nil {
		log.Fatalln("Server Error: ", err)
	}
}
