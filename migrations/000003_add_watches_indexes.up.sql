CREATE INDEX IF NOT EXISTS watches_brand_idx ON watches
    USING GIN (to_tsvector('simple', brand));
CREATE INDEX IF NOT EXISTS watches_model_idx ON watches
    USING GIN (to_tsvector('simple', model));
CREATE INDEX IF NOT EXISTS watches_dial_color_idx ON watches
    USING GIN (to_tsvector('simple', dial_color));
CREATE INDEX IF NOT EXISTS watches_strap_type_idx ON watches
    USING GIN (to_tsvector('simple', strap_type));