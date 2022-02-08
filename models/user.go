package models

type User struct {
	Id uint `json:"id" gorm:"unique" gorm:"integer" gorm:"bigserial"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique" gorm:"text"`
	Password []byte `json:"-"`
}
