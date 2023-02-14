package repository

import (
	"context"

	"github.com/muhangga/internal/entity"
)

func (r *petRepository) Save(petReq entity.Pet) (entity.Pet, error) {
	sqlStmt := `INSERT INTO pets (user_id, pet_name, gender, date_of_birth, neureted, vaccinated, friendly_with_dogs, friendly_with_cats, friendly_with_kids, microchipped, purebred, pet_nursery_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	_, err := r.db.ExecContext(ctx, sqlStmt, petReq.UserID, petReq.PetName, petReq.Gender, petReq.DateOfBirth, petReq.Neureted, petReq.Vaccinated, petReq.FriendlyWithDogs, petReq.FriendlyWithCats, petReq.FriendlyWithKidsOver10, petReq.Microchipped, petReq.Purebred, petReq.PetNurseryName)

	if err != nil {
		return petReq, err
	}

	sqlStmt = `SELECT * FROM pets`

	row := r.db.QueryRowContext(ctx, sqlStmt)
	err = row.Scan(
		&petReq.ID,
		&petReq.UserID,
		&petReq.PetName,
		&petReq.Gender,
		&petReq.DateOfBirth,
		&petReq.Neureted,
		&petReq.Vaccinated,
		&petReq.FriendlyWithDogs,
		&petReq.FriendlyWithCats,
		&petReq.FriendlyWithKidsOver10,
		&petReq.Microchipped,
		&petReq.Purebred,
		&petReq.PetNurseryName,
		&petReq.CreatedAt,
		&petReq.UpdatedAt,
	)

	if err != nil {
		return petReq, err
	}

	return petReq, nil
}

func (r *petRepository) GetAll() ([]entity.Pet, error) {
	sqlStmt := `SELECT * FROM pets`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []entity.Pet
	for rows.Next() {
		var p entity.Pet
		if err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.PetName,
			&p.Gender,
			&p.DateOfBirth,
			&p.Neureted,
			&p.Vaccinated,
			&p.FriendlyWithDogs,
			&p.FriendlyWithCats,
			&p.FriendlyWithKidsOver10,
			&p.Microchipped,
			&p.Purebred,
			&p.PetNurseryName,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		pets = append(pets, p)
	}
	return pets, nil
}

func (r *petRepository) FindPetByID(id int64) (entity.Pet, error) {
	sqlStmt := `SELECT * FROM pets WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	var p entity.Pet

	row := r.db.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(
		&p.ID,
		&p.UserID,
		&p.PetName,
		&p.Gender,
		&p.DateOfBirth,
		&p.Neureted,
		&p.Vaccinated,
		&p.FriendlyWithDogs,
		&p.FriendlyWithCats,
		&p.FriendlyWithKidsOver10,
		&p.Microchipped,
		&p.Purebred,
		&p.PetNurseryName,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return p, err
	}

	return p, nil
}
