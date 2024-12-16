package models

type user struct {
	Id unit "gorm: primarykey;auto-increment"
	Username string "jason: "username" gorm:"unique"'
	Password string "jason: "password"
	Email    string "jason: "email"gorm:"unique""
}
