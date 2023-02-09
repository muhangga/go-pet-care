package entity

type Booking struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	DoctorID    int64  `json:"doctor_id" gorm:"type:bigint;not null"`
	BookingDate string `json:"booking_date" gorm:"type:varchar(255);not null"`
	BookingTime string `json:"booking_time" gorm:"type:varchar(255);not null"`
}
