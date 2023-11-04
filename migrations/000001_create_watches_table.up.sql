CREATE TABLE IF NOT EXISTS watches (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    brand TEXT NULL,
    model TEXT NULL,
    dial_color TEXT NOT NULL,
    strap_type TEXT NOT NULL,
    diameter INT NOT NULL,
    energy TEXT NOT NULL,
    gender TEXT NOT NULL,
    price FLOAT NOT NULL,
    image_url TEXT NOT NULL,
    version INT NOT NULL DEFAULT 1
);

GRANT ALL PRIVILEGES ON watches TO watch_admin;
GRANT ALL PRIVILEGES ON SCHEMA public TO watch_admin;
GRANT ALL PRIVILEGES ON SEQUENCE watches_id_seq TO watch_admin;
ALTER DATABASE watch_database OWNER TO watch_admin;
ALTER TABLE watches OWNER TO watch_admin;