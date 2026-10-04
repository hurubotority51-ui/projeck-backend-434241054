CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    rental_id INT NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    payment_method VARCHAR(30) NOT NULL,
    payment_status VARCHAR(20) NOT NULL DEFAULT 'pending',
    paid_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_payment_rental
        FOREIGN KEY (rental_id)
        REFERENCES rentals(id)
        ON DELETE RESTRICT,

    CONSTRAINT check_payment_amount
        CHECK (amount > 0)
);