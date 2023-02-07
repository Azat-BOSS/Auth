package main

import (
	"Auth/DB"
	"Auth/Routers"
)

func main() {
	DB.ConnectDB()
	Routers.RoutingUser()
}
