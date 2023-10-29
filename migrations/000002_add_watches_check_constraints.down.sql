ALTER TABLE watches DROP CONSTRAINT IF EXISTS watches_male_diameter_check;
ALTER TABLE watches DROP CONSTRAINT IF EXISTS watches_female_diameter_check;

ALTER TABLE watches DROP CONSTRAINT IF EXISTS watches_gender_check;
ALTER TABLE watches DROP CONSTRAINT IF EXISTS watches_price_check;