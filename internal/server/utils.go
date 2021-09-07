package server

import (
	"errors"
	"fmt"
	"log"
	"os"
	"ourtool/internal/db"

	"golang.org/x/crypto/bcrypt"
)

type CustomError struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

type SuccessResponse struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

// SetOutputRef : Setting custom logging files.
func SetOutputRef(o *os.File) {
	OutputRef = o
}

func checkError(err error) *CustomError {
	var customErr *CustomError = nil
	if err != nil {
		log.Println(err.Error())
		customErr = &CustomError{Result: "fail", Msg: err.Error()}
	}
	return customErr
}

func ValidateUserDetails(user db.User) *CustomError {

	if len(user.Name) == 0 {
		if e := checkError(errors.New("please provide a name")); e != nil {
			return e
		}
	}

	if len(user.Username) == 0 {
		if e := checkError(fmt.Errorf("please provide a username")); e != nil {
			return e
		}
	}

	if len(user.Email) == 0 {
		if e := checkError(fmt.Errorf("please provide a email")); e != nil {
			return e
		}
	}

	if len(user.Password) == 0 {
		if e := checkError(fmt.Errorf("please provide a password")); e != nil {
			return e
		}
	}
	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
