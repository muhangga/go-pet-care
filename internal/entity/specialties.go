package entity

type Specialties struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
}
