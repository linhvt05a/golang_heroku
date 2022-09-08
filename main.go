package main

import (
	"example/kenva-be/db"
	"example/kenva-be/routers"

	_ "github.com/lib/pq"
)

func main() {
	db.ConnectDB()
	routers.Router()
}
