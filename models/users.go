package models

import (
	"os"
	"strings"

	u "github.com/skantuz/goreson/utils"

	"github.com/skantuz/gorm"
	"github.com/skantuz/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user us
type User struct {
	gorm.Model
	Username string `bson:"username";`
	Password string `bson:"password"`
	Email    string `bson:"email"`
	Name     string `bson:"name"`
	Lastname string `bson:"lastname"`

	Token string `bson:"token";sql:"-"`
}

//Validate incoming user details...
func (us *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(us.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(us.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", us.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (us *User) Create() map[string]interface{} {

	if resp, ok := us.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	us.Password = string(hashedPassword)

	GetDB().Create(us)

	if us.ID <= 0 {
		return u.Message(false, "Failed to create us, connection error.")
	}

	//Create new JWT token for the newly registered us
	tk := &Token{UserId: us.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	us.Token = tokenString

	us.Password = "" //delete password

	response := u.Message(true, "User has been created")
	response["us"] = us
	return response
}

func Login(email, password string) map[string]interface{} {

	us := &User{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(us).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	us.Password = ""

	//Create JWT token
	tk := &Token{UserId: us.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	us.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["us"] = us
	return resp
}

func GetUser(u uint) *User {

	acc := &User{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
