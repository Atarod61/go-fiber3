package models

type user struct {
	Id unit "jason: "id" gorm: primarykey;auto-increment"
	Username string "jason: "username" gorm:"unique"'
	Password string "jason: "password" gorm:"type:varchar(10)"
	Email    string "jason: "email"gorm:"unique""
}
