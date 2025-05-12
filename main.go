package main

import (
	"contacts-pay/config"
	"contacts-pay/routes"
)

func main() {
	config.InitDatabase()
	config.InitRedis()
	config.InitDynamoDB()

	r := routes.SetupRouter()
	r.Run(":8080")
}
