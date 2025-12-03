CREATE TABLE batch_inputs (
    sl_no SERIAL PRIMARY KEY,
    start_time TIMESTAMP,
    raw_material_name VARCHAR(50),
    raw_material_quantity INT,
    temperature INT,
    end_time TIMESTAMP
);

SELECT 'Database initialized successfullyyyy!' as status;