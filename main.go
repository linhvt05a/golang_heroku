package main

import (
	"fmt"
	"os"
)

// "example/kenva-be/db"
// "example/kenva-be/routers"

func main() {
	KEY := os.Getenv("TEST")
	fmt.Println("key", KEY)
	// db.ConnectDB()
	// routers.Router()

}
