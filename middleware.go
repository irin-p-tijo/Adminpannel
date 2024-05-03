package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Validatecookies(c *gin.Context) bool {
	cookies, _ := c.Cookie("cookies")
	if cookies == "" {
		fmt.Println("error in creating cookies")
		return false
	} else {
		fmt.Println("cookies", cookies)
		return true
	}
}
func deletecookies(c *gin.Context) {
	c.SetCookie("cookies", "", "0", "", false, true)
	fmt.Println("the cookie is deleted")
}
func Findrole(c *gin.Context) {
	cookies, _ := c.Cookie("cookies")
	if cookies == "" {
		fmt.Println("cookie is not found")
	}
	token, err := jwt.Parse(cookiesfunc(token * jwt.Token))
}
