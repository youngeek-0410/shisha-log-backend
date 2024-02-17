BEGIN;

INSERT INTO
    flavor_brands name
VALUES
    "Al Fakher";

SET
    flavor_brand_id = LAST_INSERT_ID ();

INSERT INTO
    flavors (brand_id, name)
VALUES
    (flavor_brand_id, "ローズ");

INSERT INTO
    flavors (brand_id, name)
VALUES
    (flavor_brand_id, "ピーチ");

INSERT INTO
    flavors (brand_id, name)
VALUES
    (flavor_brand_id, "カルダモン");

COMMIT;

-- INSERT INTO
--     flavor_brands (id, name)
-- VALUES
--     (UUID_TO_BIN ("", 1), "Afzal",);
-- INSERT INTO
--     flavor_brands (id, name)
-- VALUES
--     (UUID_TO_BIN ("", 1), "Dozaj",);