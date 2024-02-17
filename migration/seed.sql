SET @flavor_brand_id_1 = UUID ();
SET @flavor_brand_id_2 = UUID ();
SET @flavor_brand_id_3 = UUID ();

SET @flavor_id_1 = UUID ();
SET @flavor_id_2 = UUID ();
SET @flavor_id_3 = UUID ();
SET @flavor_id_4 = UUID ();
SET @flavor_id_5 = UUID ();
SET @flavor_id_6 = UUID ();

-- flavor_brands
INSERT INTO
    flavor_brands (id, name)
VALUES
    (UUID_TO_BIN (@flavor_brand_id_1, 1), 'Al Fakher');

INSERT INTO
    flavor_brands (id, name)
VALUES
    (UUID_TO_BIN (@flavor_brand_id_2, 1), 'Afzal');

INSERT INTO
    flavor_brands (id, name)
VALUES
    (UUID_TO_BIN (@flavor_brand_id_3, 1), 'Dozaj');

-- flavors
INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@flavor_id_1, 1), UUID_TO_BIN (@flavor_brand_id_1, 1), 'ローズ');

INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@flavor_id_2, 1), UUID_TO_BIN (@flavor_brand_id_1, 1), 'ピーチ');

INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@flavor_id_3, 1), UUID_TO_BIN (@flavor_brand_id_1, 1), 'カルダモン');

INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@flavor_id_4, 1), UUID_TO_BIN (@flavor_brand_id_2, 1), 'アールグレイ');

INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@flavor_id_5, 1), UUID_TO_BIN (@flavor_brand_id_2, 1), 'ジンジャーエール');

INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@flavor_id_6, 1), UUID_TO_BIN (@flavor_brand_id_3, 1), 'ブラックティー');
