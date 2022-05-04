package model

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	UserID       int64        `db:"user_id"`
	Username     string       `db:"username"`
	Password     string       `db:"password"` //salted
	DateAdded    sql.NullTime `db:"date_added"`
	DateModified sql.NullTime `db:"date_modified"`
}

func (u *User) SetPassword(password string) {
	u.Password = HashPassword(password)
}

type Api struct {
	UserID       int64        `db:"user_id"`
	ApiKey       int          `db:"api_key"`
	DateAdded    sql.NullTime `db:"date_added"`
	DateModified sql.NullTime `db:"date_modified"`
}

type Transaction struct {
	UserID       int64        `db:"user_id"`
	Action       string       `db:"action"`
	DateAdded    sql.NullTime `db:"date_added"`
	DateModified sql.NullTime `db:"date_modified"`
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
