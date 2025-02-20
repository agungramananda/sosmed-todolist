CREATE TABLE brands (
    brand_id SERIAL PRIMARY KEY,
    brand VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE platforms (
    platform_id SERIAL PRIMARY KEY,
    platform VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TYPE task_status AS ENUM('Pending', 'Completed', 'Scheduled');

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    brand_id INT REFERENCES brands(brand_id) ON DELETE CASCADE,
    platform_id INT REFERENCES platforms(platform_id) ON DELETE CASCADE,
    due_date TIMESTAMP NOT NULL,
    payment DECIMAL(10,2) DEFAULT 0.00,
    status task_status NOT NULL DEFAULT 'Pending', 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);