package entity

type Additional struct {
	ID               int64 `json:"id" gorm:"primary_key"`
	PetID            int64 `json:"pet_id" gorm:"type:bigint;not null"`
	Neureted         bool  `json:"neureted" gorm:"type:bool;not null"`
	Vaccinated       bool  `json:"vaccinated" gorm:"type:bool;not null"`
	FriendlyWithDogs bool  `json:"friendly_with_dogs" gorm:"type:bool;not null"`
	FriendlyWithCats bool  `json:"friendly_with_cats" gorm:"type:bool;not null"`

	FriendlyWithKidsOver10 bool   `json:"friendly_with_kids_over_10" gorm:"type:bool;not null"`
	Microchipped           bool   `json:"microchipped" gorm:"type:bool;not null"`
	Purebred               bool   `json:"purebred" gorm:"type:bool;not null"`
	PetNurseryName         string `json:"pet_nursery_name" gorm:"type:text;not null"`
}
