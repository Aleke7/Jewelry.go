-- ALTER TABLE watches ADD CONSTRAINT watches_male_diameter_check
--     CHECK ( (lower(gender) = 'male' OR (lower(gender) = 'men'
--                                              OR (lower(gender) = 'm'))
--     AND diameter >= 38 AND diameter <= 46));
--
-- ALTER TABLE watches ADD CONSTRAINT watches_female_diameter_check
--     CHECK ( (lower(gender) = 'female' OR (lower(gender) = 'women'
--                                                OR (lower(gender) = 'f'))
--         AND diameter >= 26 AND diameter <= 36));

ALTER TABLE watches ADD CONSTRAINT watches_gender_check
    CHECK ( (lower(gender) = 'male' OR lower(gender) = 'men' OR lower(gender) = 'm')
                OR (lower(gender) = 'female' OR lower(gender) = 'women' OR lower(gender) = 'f'));

ALTER TABLE watches ADD CONSTRAINT watches_price_check
    CHECK ( price > 0 )