//Controllers/User.go

package Controllers

import (
 "sqltryout/Models"
 "fmt"
 "net/http"
 "github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
 var user []Models.User
 err := Models.GetAllUsers(&user)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, user)
 }
}

func RegisterUser(c *gin.Context){
	var user Models.User
	c.BindJSON(&user)
	username := user.Name
	password := user.Password
	err := Models.CreateUser(&user, username, password)
	if err!= nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.Set("is_logged_in", true)
		c.JSON(http.StatusOK, user)
	}
}

//userLogin ... Log user in
func UserLogin(c *gin.Context) {
 var user Models.User
 username := user.Name
 password := user.Password
 err := Models.UserValid(username, password)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.Set("is_logged_in", true)
  c.JSON(http.StatusOK, user)
 }
}

func UserLogout(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/user-api")
}

//UpdateUser ... Update the user information
// func UpdateUser(c *gin.Context) {
//  var user Models.User
//  id := c.Params.ByName("id")
//  err := Models.GetUserByID(&user, id)
//  if err != nil {
//   c.JSON(http.StatusNotFound, user)
//  }
//  c.BindJSON(&user)
//  err = Models.UpdateUser(&user, id)
//  if err != nil {
//   c.AbortWithStatus(http.StatusNotFound)
//  } else {
//   c.JSON(http.StatusOK, user)
//  }
// }

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
 var user Models.User
 id := c.Params.ByName("id")
 err := Models.DeleteUser(&user, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
 }
}