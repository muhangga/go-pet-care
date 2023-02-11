CREATE TABLE IF NOT EXISTS additionals (
    id SERIAL PRIMARY KEY,
    pet_id BIGINT NOT NULL,
    neureted BOOLEAN NOT NULL,
    vaccinated BOOLEAN NOT NULL,
    friendly_with_dogs BOOLEAN NOT NULL,
    friendly_with_cats BOOLEAN NOT NULL,
    friendly_with_kids BOOLEAN NOT NULL,
    microchipped BOOLEAN NOT NULL,
    purebred BOOLEAN NOT NULL,
    pet_nursery_name VARCHAR(255) NOT NULL
);