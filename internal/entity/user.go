package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID          sql.NullInt64  `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName    string         `json:"full_name" gorm:"type:varchar(255);not null"`
	Email       string         `json:"email" gorm:"type:varchar(255);not null"`
	Password    string         `json:"password" gorm:"type:varchar(255);not null"`
	PhoneNumber sql.NullString `json:"phone_number" gorm:"type:varchar(255);nullString"`
	Gender      sql.NullString `json:"gender" gorm:"type:varchar(255);null"`
	Avatar      sql.NullString `json:"avatar" gorm:"type:varchar(255);null"`
	AboutMe     sql.NullString `json:"about_me" gorm:"type:varchar(255);null"`
	Role        string         `json:"role" gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"type:timestamp;not null"`
}
