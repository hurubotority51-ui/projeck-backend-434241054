CREATE TABLE vehicles (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    brand VARCHAR(50) NOT NULL,
    model VARCHAR(50),
    license_plate VARCHAR(20) UNIQUE NOT NULL,
    year INT,
    price_per_day NUMERIC(12,2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'available',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_vehicle_category
        FOREIGN KEY (category_id)
        REFERENCES vehicle_categories(id)
        ON DELETE RESTRICT
);