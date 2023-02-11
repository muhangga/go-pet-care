CREATE TABLE IF NOT EXISTS booking (
    id bigserial PRIMARY KEY,
    doctor_id INT NOT NULL,
    booking_date DATE NOT NULL,
    booking_time TIME NOT NULL,
    status VARCHAR(255) NOT NULL
)