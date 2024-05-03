package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/middleware"
)

func signuphandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	ok := middleware.Validatecookies(c)
	if !ok {

		c.HTML(http.StatusOK, "signup.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/")

}
func SignupPost(c *gin.Context) {
	c.Header("cache-control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	var errors models.error
	username := c.Request.FormValue("Name")
	useremail := c.Request.FormValue("Email")
	if username == "" {
		errors.NameError = "Email should not be empty"
		c.HTML(http.StatusBadRequest, "signup.html", errors)
		return
	}
	if useremail == "" {
		errors.EmailError = "email should not be empty"
		c.HTML(http.StatusBadRequest, "signup.html", errors)
		return
	}
	password := c.Request.FormValue("Password")
	confirmpassword := c.Request.FormValue("confirmpassword")
	if password == "" {
		errors.PasswordError = "password does not be empty"
		c.HTML(http.StatusBadRequest, "signup.html", errors)
	}
	if password != confirmpassword {
		errors.PasswordError = "password does not match"
		c.HTML(http.StatusBadRequest, "signup.html", errors)
		return
	}
}
func loginhandler(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	ok := middleware.Validatecookies(c)
	role, _, _ := middleware.Findrole(c)
	if !ok {

		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		if role == "user" {
			c.Redirect(http.StatusFound, "/home")
			return
		} else if role == "admin" {
			c.Redirect(http.StatusBadRequest, "/admin")
			return
		}
		c.HTML(http.StatusBadRequest, "login.html", nil)
	}
}
func loginpost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	var error models.errors
	NewEmail := c.Request.FormValue("Email")
	Newpassword := c.Request.FormValue("Password")
	var compare models.compare
	if err:=db,Db,Raw("")
}
