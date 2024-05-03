package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("do not load the .env files")
	}
	router := gin.New()
	router.GET("/", handlers.loginHandler)
	router.POST("/", handlers.loginpost)
	router.GET("/signup", handlers.Signuphandler)
	router.POST("/signup", handlers.signuppost)
	router.GET("/home", handlers.homehandler)
	router.GET("/logout", handlers.logouthandler)
}
