//Models/User.go

package Models

import (
	"errors"
	"strings"
 "sqltryout/Config"
 "fmt"
 _ "github.com/go-sql-driver/mysql"
)

type user struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

var userList = []user{}

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
 if err = Config.DB.Find(user).Error; err != nil {
  return err
 }
 return nil
}

//CreateUser ... Insert New data
func CreateUser(userv *User, username string, password string) (error) {
 if strings.TrimSpace(password) == "" {
        return errors.New("The password can't be empty")
 } else if !isUsernameAvailable(username) {
        return errors.New("Username is not available")
 } 
 
 u := user{Username: username, Password: password}

 userList = append(userList, u)

 if err := Config.DB.Create(userv).Error; err != nil {
  return errors.New("Unable to upload to server")
 }
 return nil
}

func isUsernameAvailable(username string) bool {
    for _, element := range userList {
        if element.Username == username {
            return false
        }
    }
    return true
}

//UserValid...checks if user is registered
func UserValid(username string, password string) (error) {
 for _, element := range userList {
	 if username == element.Username && password == element.Password {
		return nil
	 }
 }
 return errors.New("Username or Password incorrect")
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
 fmt.Println(user)
 Config.DB.Save(user)
 return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
 Config.DB.Where("id = ?", id).Delete(user)
 return nil
}