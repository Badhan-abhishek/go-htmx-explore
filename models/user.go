package models

type User struct {
	Model
	Name     string  `json:"name"`
	Email    *string `json:"email" gorm:"unique"`
	Password []byte  `json:"-"`
}
