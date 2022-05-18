//Models/UserModel.go
package Models

type User struct {
 Id      uint   `json:"id"`
 Name    string `json:"name" binding:"required"`
 Password string `json:"password" binding:"required"`
 Email   string `json:"email"`
 Address string `json:"address"`
}

func (b *User) TableName() string {
 return "user"
}