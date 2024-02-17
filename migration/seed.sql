SET @flavor = "1f850241-1c1f-f205-a797-da0b9f459284";
INSERT INTO
    flavor_brands (id, name)
VALUES
    (UUID_TO_BIN (@flavor, 1), 'Al Fakher');

INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (UUID_TO_BIN ("db62c51e-7355-7a2b-4f24-1275465eb278", 1), UUID_TO_BIN (@flavor, 1), 'ローズ');