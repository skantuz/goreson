package models

import (
	u "github.com/skantuz/goreson/utils"
	"github.com/skantuz/gorm"
)

//a struct to rep Role
type Role struct {
	gorm.Model
	Name  string `bson:"name"`
	Users []User `gorm:"foreignkey:RoleId"`
}

//Validate incoming user details...

func (role *Role) Create() map[string]interface{} {

	GetDB().Create(role)

	response := u.Message(true, "Role has been created")
	response["role"] = role
	return response
}

func GetRole(u uint) *Role {

	acc := &Role{}
	GetDB().Table("roles").Where("id = ?", u).First(acc)
	if acc.Name == "" { //Role not found!
		return nil
	}
	return acc
}
