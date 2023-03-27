package main

import (
	"fmt"

	"github.com/JeerasakTH/template-go-postgres/database"
	"github.com/JeerasakTH/template-go-postgres/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("ConnectDB error:")
		panic(err)
	}
	r := gin.Default()

	router.Router(r, db)

	r.Run()
}
