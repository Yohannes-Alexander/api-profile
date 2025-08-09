package main

import (
	"fmt"
	"log"

	"github.com/Yohannes-Alexander/api-profile/config"
	"github.com/Yohannes-Alexander/api-profile/internal/router"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.SetupRouter(db)
	fmt.Println("Server running at :8080")
	r.Run(":8080")
}
