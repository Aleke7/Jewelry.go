CREATE TABLE IF NOT EXISTS watches (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    brand text NULL,
    model text NULL,
    dial_color text NOT NULL,
    strap_type text NOT NULL,
    diameter integer NOT NULL,
    energy text NOT NULL,
    gender text NOT NULL,
    price float NOT NULL,
    image_url text NOT NULL,
    version integer NOT NULL DEFAULT 1
);

GRANT ALL PRIVILEGES ON watches TO watch_admin;
GRANT ALL PRIVILEGES ON SCHEMA public TO watch_admin;
GRANT ALL PRIVILEGES ON SEQUENCE watches_id_seq TO watch_admin;
ALTER DATABASE watch_database OWNER TO watch_admin;
ALTER TABLE watches OWNER TO watch_admin;