//Routes/Routes.go
package Routes

import (
 "sqltryout/Controllers"
 "github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes

func SetupRouter() *gin.Engine {
 r := gin.Default()
 grp1 := r.Group("/user-api")
 {
  grp1.GET("user", Controllers.GetUsers)
  grp1.POST("user/register", Controllers.RegisterUser)
  grp1.POST("user/login", Controllers.UserLogin)
  grp1.POST("user/logout", Controllers.UserLogout)
  //grp1.PUT("user/:id", Controllers.UpdateUser)
  grp1.DELETE("user/:id", Controllers.DeleteUser)
 }
 return r
}