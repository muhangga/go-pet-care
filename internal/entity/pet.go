package entity

import "time"

type Pet struct {
	ID                     int64     `json:"id"`
	UserID                 int64     `json:"user_id"`
	PetName                string    `json:"pet_name"`
	Gender                 string    `json:"gender"`
	DateOfBirth            time.Time `json:"date_of_birth"`
	Neureted               bool      `json:"neureted"`
	Vaccinated             bool      `json:"vaccinated"`
	FriendlyWithDogs       bool      `json:"friendly_with_dogs"`
	FriendlyWithCats       bool      `json:"friendly_with_cats" `
	FriendlyWithKidsOver10 bool      `json:"friendly_with_kids_over_10"`
	Microchipped           bool      `json:"microchipped"`
	Purebred               bool      `json:"purebred"`
	PetNurseryName         string    `json:"pet_nursery_name"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type PetRequest struct {
	PetName                string    `json:"pet_name"`
	Gender                 string    `json:"gender"`
	DateOfBirth            time.Time `json:"date_of_birth"`
	Neureted               bool      `json:"neureted"`
	Vaccinated             bool      `json:"vaccinated"`
	FriendlyWithDogs       bool      `json:"friendly_with_dogs"`
	FriendlyWithCats       bool      `json:"friendly_with_cats" `
	FriendlyWithKidsOver10 bool      `json:"friendly_with_kids_over_10"`
	Microchipped           bool      `json:"microchipped"`
	Purebred               bool      `json:"purebred"`
	PetNurseryName         string    `json:"pet_nursery_name"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
