package db

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Key       string    `json:"key"`
	Password  string    `json:"password"`
}

func (u User) Prepare(data io.ReadCloser) (User, *CustomError) {

	json.NewDecoder(data).Decode(&u)

	err := ValidateUserDetails(u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (u User) Before() (User, *CustomError) {

	hashedPassword, err := Hash(u.Password)
	if e := checkError(err); err != nil {
		return User{}, e
	}
	u.Password = string(hashedPassword)
	u.Key = fmt.Sprintf("%d_%d", time.Now().UnixNano(), rand.Int())

	return u, nil
}

func (u User) Save(db *gorm.DB) (User, *CustomError) {
	err := db.Debug().Create(&u).Error
	if e := checkError(err); err != nil {
		return User{}, e
	}
	return u, nil
}
