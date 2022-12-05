package main

import (
	"example/kenva-be/db"
	"example/kenva-be/routers"
)
func main() {
	db.ConnectDB()
	routers.Router()

}
