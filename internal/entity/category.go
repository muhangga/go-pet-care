package entity

type Category struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Icon string `json:"icon" gorm:"type:varchar(255);null"`
}
