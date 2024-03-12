CREATE TABLE IF NOT EXISTS pin_pull (
    id SERIAL PRIMARY KEY,
    chat_id BIGINT,
    message_id INT,
    created_at DATE,
    type_pull VARCHAR(255)
);