-- flavor_brands
SET @flavor_brand_id_1 = UUID ();
SET @flavor_brand_id_2 = UUID ();
SET @flavor_brand_id_3 = UUID ();

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
SET @flavor_id_1 = UUID ();
SET @flavor_id_2 = UUID ();
SET @flavor_id_3 = UUID ();
SET @flavor_id_4 = UUID ();
SET @flavor_id_5 = UUID ();
SET @flavor_id_6 = UUID ();

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

-- bowl_brands
SET @bowl_brand_id_1 = UUID ();
SET @bowl_brand_id_2 = UUID ();
SET @bowl_brand_id_3 = UUID ();

INSERT INTO
    bowl_brands (id, name)
VALUES
    (UUID_TO_BIN (@bowl_brand_id_1, 1), 'Hookah John');

INSERT INTO
    bowl_brands (id, name)
VALUES
    (UUID_TO_BIN (@bowl_brand_id_2, 1), 'Moon Hookah');

INSERT INTO
    bowl_brands (id, name)
VALUES
    (UUID_TO_BIN (@bowl_brand_id_3, 1), 'Cosmo bowl');

-- bowls
SET @bowl_id_1 = UUID ();
SET @bowl_id_2 = UUID ();
SET @bowl_id_3 = UUID ();

INSERT INTO
    bowls (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@bowl_id_1, 1), UUID_TO_BIN (@bowl_brand_id_1, 1), '80feet ESPANA');

INSERT INTO
    bowls (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@bowl_id_2, 1), UUID_TO_BIN (@bowl_brand_id_2, 1), 'Moon Killer 2.0');

INSERT INTO
    bowls (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@bowl_id_3, 1), UUID_TO_BIN (@bowl_brand_id_3, 1), 'Turkish Predator');

-- charcoal_brands
SET @charcoal_brand_id_1 = UUID ();
SET @charcoal_brand_id_2 = UUID ();

INSERT INTO
    charcoal_brands (id, name)
VALUES
    (UUID_TO_BIN (@charcoal_brand_id_1, 1), 'KINGCO');

INSERT INTO
    charcoal_brands (id, name)
VALUES
    (UUID_TO_BIN (@charcoal_brand_id_2, 1), 'COCOMELT');

-- charcoals
SET @charcoal_id_1 = UUID ();
SET @charcoal_id_2 = UUID ();
SET @charcoal_id_3 = UUID ();

INSERT INTO
    charcoals (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@charcoal_id_1, 1), UUID_TO_BIN (@charcoal_brand_id_1, 1), 'Coconuts flat cube');

INSERT INTO
    charcoals (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@charcoal_id_2, 1), UUID_TO_BIN (@charcoal_brand_id_2, 1), 'Coconuts flat cube 22');

INSERT INTO
    charcoals (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@charcoal_id_3, 1), UUID_TO_BIN (@charcoal_brand_id_2, 1), 'Coconuts flat cube 26');

-- heat_management_brands
SET @heat_management_brand_id_1 = UUID ();
SET @heat_management_brand_id_2 = UUID ();

INSERT INTO
    heat_management_brands (id, name)
VALUES
    (UUID_TO_BIN (@heat_management_brand_id_1, 1), 'YIMI');

INSERT INTO
    heat_management_brands (id, name)
VALUES
    (UUID_TO_BIN (@heat_management_brand_id_2, 1), 'ONMO');

-- heat_managements
SET @heat_management_id_1 = UUID ();
SET @heat_management_id_2 = UUID ();

INSERT INTO
    heat_managements (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@heat_management_id_1, 1), UUID_TO_BIN (@heat_management_brand_id_1, 1), 'ターキッシュリッド');

INSERT INTO
    heat_managements (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@heat_management_id_2, 1), UUID_TO_BIN (@heat_management_brand_id_2, 1), 'HMD - Black');

-- bottle_brands
SET @bottle_brand_id_1 = UUID ();
SET @bottle_brand_id_2 = UUID ();

INSERT INTO
    bottle_brands (id, name)
VALUES
    (UUID_TO_BIN (@bottle_brand_id_1, 1), 'Stimulation');

INSERT INTO
    bottle_brands (id, name)
VALUES
    (UUID_TO_BIN (@bottle_brand_id_2, 1), 'Cloud one mini');

-- bottles
SET @bottle_id_1 = UUID ();
SET @bottle_id_2 = UUID ();

INSERT INTO
    bottles (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@bottle_id_1, 1), UUID_TO_BIN (@bottle_brand_id_1, 1), 'Ultimate one pro');

INSERT INTO
    bottles (id, brand_id, name)
VALUES
    (UUID_TO_BIN (@bottle_id_2, 1), UUID_TO_BIN (@bottle_brand_id_2, 1), 'Shisha Backs');

-- users
SET @user_id_1 = UUID ();

INSERT INTO
    users (id, name)
VALUES
    (UUID_TO_BIN (@user_id_1, 1), 'Ryusei Ito');

-- user_flavors
SET @user_flavor_id_1 = UUID ();
SET @user_flavor_id_2 = UUID ();
SET @user_flavor_id_3 = UUID ();
SET @user_flavor_id_4 = UUID ();
SET @user_flavor_id_5 = UUID ();
SET @user_flavor_id_6 = UUID ();

INSERT INTO
    user_flavors (id, flavor_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_flavor_id_1, 1),
        UUID_TO_BIN (@flavor_id_1, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_flavors (id, flavor_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_flavor_id_2, 1),
        UUID_TO_BIN (@flavor_id_2, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_flavors (id, flavor_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_flavor_id_3, 1),
        UUID_TO_BIN (@flavor_id_3, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_flavors (id, flavor_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_flavor_id_4, 1),
        UUID_TO_BIN (@flavor_id_4, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_flavors (id, flavor_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_flavor_id_5, 1),
        UUID_TO_BIN (@flavor_id_5, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_flavors (id, flavor_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_flavor_id_6, 1),
        UUID_TO_BIN (@flavor_id_6, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

-- user_bowls
SET @user_bowl_id_1 = UUID ();
SET @user_bowl_id_2 = UUID ();
SET @user_bowl_id_3 = UUID ();

INSERT INTO
    user_bowls (id, bowl_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_bowl_id_1, 1),
        UUID_TO_BIN (@bowl_id_1, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_bowls (id, bowl_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_bowl_id_2, 1),
        UUID_TO_BIN (@bowl_id_2, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_bowls (id, bowl_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_bowl_id_3, 1),
        UUID_TO_BIN (@bowl_id_3, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

-- user_charcoals
SET @user_charcoal_id_1 = UUID ();
SET @user_charcoal_id_2 = UUID ();
SET @user_charcoal_id_3 = UUID ();

INSERT INTO
    user_charcoals (id, charcoal_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_charcoal_id_1, 1),
        UUID_TO_BIN (@charcoal_id_1, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_charcoals (id, charcoal_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_charcoal_id_2, 1),
        UUID_TO_BIN (@charcoal_id_2, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_charcoals (id, charcoal_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_charcoal_id_3, 1),
        UUID_TO_BIN (@charcoal_id_3, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

-- user_heat_managements
SET @user_heat_management_id_1 = UUID ();
SET @user_heat_management_id_2 = UUID ();

INSERT INTO
    user_heat_managements (id, heat_management_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_heat_management_id_1, 1),
        UUID_TO_BIN (@heat_management_id_1, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_heat_managements (id, heat_management_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_heat_management_id_2, 1),
        UUID_TO_BIN (@heat_management_id_2, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

-- user_bottles
SET @user_bottle_id_1 = UUID ();
SET @user_bottle_id_2 = UUID ();

INSERT INTO
    user_bottles (id, bottle_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_bottle_id_1, 1),
        UUID_TO_BIN (@bottle_id_1, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );

INSERT INTO
    user_bottles (id, bottle_id, user_id)
VALUES
    (
        UUID_TO_BIN (@user_bottle_id_2, 1),
        UUID_TO_BIN (@bottle_id_2, 1),
        UUID_TO_BIN (@user_id_1, 1)
    );
