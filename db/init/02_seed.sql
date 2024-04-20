-- flavor_brands
\set FLAVOR_BRAND_ID_1 69d7e1bb-ca48-4a5d-b40f-2f0213761b30
\set FLAVOR_BRAND_ID_2 f574dcf9-fb93-46b0-98cd-7cd1917c3449
\set FLAVOR_BRAND_ID_3 68557d8d-016d-4550-99d9-5c6afac8b0d6

INSERT INTO 
    flavor_brands (id, name)
VALUES 
    (:'FLAVOR_BRAND_ID_1', 'Al Fakher'),
    (:'FLAVOR_BRAND_ID_2', 'Afzal'),
    (:'FLAVOR_BRAND_ID_3', 'Doza');

-- flavors
\set FLAVOR_ID_1 a243b6ab-3cd2-47ea-90fd-5492b9e3a4a7
\set FLAVOR_ID_2 5d9acd1f-1fd1-4222-91a7-d8cb06d0fd28
\set FLAVOR_ID_3 b44d6542-3eb9-47d9-b577-5d6c55dc3e13
\set FLAVOR_ID_4 89637a8e-f0da-41e5-ab42-d2e70985c8ad
\set FLAVOR_ID_5 793af5f9-0ac5-491d-9666-fbc0b597ac25
\set FLAVOR_ID_6 2f5de5db-6bbb-427f-9df3-4252d1a57357

INSERT INTO
    flavors (id, brand_id, name)
VALUES
    (:'FLAVOR_ID_1', :'FLAVOR_BRAND_ID_1', 'ローズ'),
    (:'FLAVOR_ID_2', :'FLAVOR_BRAND_ID_1', 'ピーチ'),
    (:'FLAVOR_ID_3', :'FLAVOR_BRAND_ID_1', 'カルダモン'),
    (:'FLAVOR_ID_4', :'FLAVOR_BRAND_ID_2', 'アールグレイ'),
    (:'FLAVOR_ID_5', :'FLAVOR_BRAND_ID_2', 'ジンジャーエール'),
    (:'FLAVOR_ID_6', :'FLAVOR_BRAND_ID_3', 'ブラックティー');

-- bowl_brands
\set BOWL_BRAND_ID_1 ca0cb44e-8cd7-46cb-bf74-542776c86f7e
\set BOWL_BRAND_ID_2 4ae49d87-edcc-4d83-b238-dc087ae22824
\set BOWL_BRAND_ID_3 a3736958-89ba-4c51-ac8d-478f65fdc9b9

INSERT INTO 
    bowl_brands (id, name)
VALUES 
    (:'BOWL_BRAND_ID_1', 'Hookah John'),
    (:'BOWL_BRAND_ID_2', 'Moon Hookah'),
    (:'BOWL_BRAND_ID_3', 'Cosmo bowl');

-- bowls
\set BOWL_ID_1 079bdc27-533d-4a2d-a9a6-910e4d09500b
\set BOWL_ID_2 e50b981d-19ac-4940-9ace-e653584cdbf4
\set BOWL_ID_3 f80d613d-a96e-4467-891b-0cde22979869

INSERT INTO
    bowls (id, brand_id, name)
VALUES
    (:'BOWL_ID_1', :'BOWL_BRAND_ID_1', '80feet ESPANA'),
    (:'BOWL_ID_2', :'BOWL_BRAND_ID_2', 'Moon Killer 2.0'),
    (:'BOWL_ID_3', :'BOWL_BRAND_ID_3', 'Turkish Predator');

-- charcoal_brands
\set CHARCOAL_BRAND_ID_1 1a349a5c-978d-4333-99e3-d5ac357e1948
\set CHARCOAL_BRAND_ID_2 57a546f1-b1b2-400a-8bf7-fca6a84ad331

INSERT INTO 
    charcoal_brands (id, name)
VALUES 
    (:'CHARCOAL_BRAND_ID_1', 'KINGCO'),
    (:'CHARCOAL_BRAND_ID_2', 'COCOMELT');

-- charcoals
\set CHARCOAL_ID_1 8e1771e9-8bf8-4ce8-ad0d-cba0795536f1
\set CHARCOAL_ID_2 5251f84c-34b9-445a-a6a9-7409b7279c48
\set CHARCOAL_ID_3 bf60eccd-696b-48f7-9b39-25b657b98112

INSERT INTO
    charcoals (id, brand_id, name)
VALUES
    (:'CHARCOAL_ID_1', :'CHARCOAL_BRAND_ID_1', 'Coconuts flat cube'),
    (:'CHARCOAL_ID_2', :'CHARCOAL_BRAND_ID_2', 'Coconuts flat cube 22'),
    (:'CHARCOAL_ID_3', :'CHARCOAL_BRAND_ID_2', 'Coconuts flat cube 26');

-- heat_management_brands
\set HEAT_MANAGEMENT_BRAND_ID_1 ee489ae2-317e-42cb-acbe-c2469d174c64
\set HEAT_MANAGEMENT_BRAND_ID_2 00882f21-fbcb-472f-b5c4-968a8bd1d42c

INSERT INTO
    heat_management_brands (id, name)
VALUES
    (:'HEAT_MANAGEMENT_BRAND_ID_1', 'YIMI'),
    (:'HEAT_MANAGEMENT_BRAND_ID_2', 'ONMO');

-- heat_managements
\set HEAT_MANAGEMENT_ID_1 23aee072-7869-405d-9d00-14667a52b529
\set HEAT_MANAGEMENT_ID_2 c68fc180-52da-405b-ac11-7ff9459e3592

INSERT INTO
    heat_managements (id, brand_id, name)
VALUES
    (:'HEAT_MANAGEMENT_ID_1', :'HEAT_MANAGEMENT_BRAND_ID_1', 'ターキッシュリッド'),
    (:'HEAT_MANAGEMENT_ID_2', :'HEAT_MANAGEMENT_BRAND_ID_2', 'HMD - Black');

-- bottle_brands
\set BOTTLE_BRAND_1 e6dfd7fd-3b2e-494b-9c85-3a8270b16fb2
\set BOTTLE_BRAND_2 7df5a7d5-2582-4574-8fc3-3972dd15f070

INSERT INTO
    bottle_brands (id, name)
VALUES
    (:'BOTTLE_BRAND_1', 'Stimulation'),
    (:'BOTTLE_BRAND_2', 'Cloud one mini');

-- bottles
\set BOTTLE_ID_1 d61274a7-3efa-4852-995d-9263b9c5a06e
\set BOTTLE_ID_2 2d3b1b18-57bb-4604-9a5b-ad912f972dba

INSERT INTO
    bottles (id, brand_id, name)
VALUES
    (:'BOTTLE_ID_1', :'BOTTLE_BRAND_1', 'Ultimate one pro'),
    (:'BOTTLE_ID_2', :'BOTTLE_BRAND_2', 'Shisha Backs');

-- users
\set USER_ID_1 3998ff35-311e-4b8f-88a7-f8e4cb0540e3

INSERT INTO
    users (id, name)
VALUES
    (:'USER_ID_1', 'Ryusei Ito');

-- user_flavors
\set USER_FLAVOR_ID_1 be7199fa-74ed-49da-8cbe-0b51edfabbe4
\set USER_FLAVOR_ID_2 2e906771-a850-4f4b-a023-18a2b0c5eb07
\set USER_FLAVOR_ID_3 01e0fe8f-e736-450b-8e9e-9aaddc294815
\set USER_FLAVOR_ID_4 2f3c47fa-82b3-4595-a3de-995ed21f5c5e
\set USER_FLAVOR_ID_5 c26d6c31-fe16-47c1-92a8-a73f54d072d1
\set USER_FLAVOR_ID_6 b2580f89-13bf-49eb-acdf-7f0e6bd2f762

INSERT INTO
    user_flavors (id, flavor_id, user_id)
VALUES
    (:'USER_FLAVOR_ID_1', :'FLAVOR_ID_1', :'USER_ID_1'),
    (:'USER_FLAVOR_ID_2', :'FLAVOR_ID_2', :'USER_ID_1'),
    (:'USER_FLAVOR_ID_3', :'FLAVOR_ID_3', :'USER_ID_1'),
    (:'USER_FLAVOR_ID_4', :'FLAVOR_ID_4', :'USER_ID_1'),
    (:'USER_FLAVOR_ID_5', :'FLAVOR_ID_5', :'USER_ID_1'),
    (:'USER_FLAVOR_ID_6', :'FLAVOR_ID_6', :'USER_ID_1');

-- user_bowls
\set USER_BOWL_ID_1 3f14fbbc-9941-4358-8f4b-5ad23b25cbb3
\set USER_BOWL_ID_2 3576c302-4520-4d51-b17f-763289458020
\set USER_BOWL_ID_3 65c7498a-0253-4662-9023-2f618510fcf8

INSERT INTO
    user_bowls (id, bowl_id, user_id)
VALUES
    (:'USER_BOWL_ID_1', :'BOWL_ID_1', :'USER_ID_1'),
    (:'USER_BOWL_ID_2', :'BOWL_ID_2', :'USER_ID_1'),
    (:'USER_BOWL_ID_3', :'BOWL_ID_3', :'USER_ID_1');

-- user_charcoals
\set USER_CHARCOAL_ID_1 7378174f-8ac8-4dfa-9e86-154ae3307b2c
\set USER_CHARCOAL_ID_2 207d0ac7-a71e-4ba4-99cf-d897cc5b5084
\set USER_CHARCOAL_ID_3 ef7d4feb-2b88-4668-bbed-fc63feb7a56e

INSERT INTO
    user_charcoals (id, charcoal_id, user_id)
VALUES
    (:'USER_CHARCOAL_ID_1', :'CHARCOAL_ID_1', :'USER_ID_1'),
    (:'USER_CHARCOAL_ID_2', :'CHARCOAL_ID_2', :'USER_ID_1'),
    (:'USER_CHARCOAL_ID_3', :'CHARCOAL_ID_3', :'USER_ID_1');

-- user_heat_managements
\set USER_HEAT_MANAGEMENT_ID_1 dd555659-dcf8-4555-84b0-2b864819213b
\set USER_HEAT_MANAGEMENT_ID_2 a8263fae-78ca-4326-9a02-7ec08d8f7d0e

INSERT INTO
    user_heat_managements (id, heat_management_id, user_id)
VALUES
    (:'USER_HEAT_MANAGEMENT_ID_1', :'HEAT_MANAGEMENT_ID_1', :'USER_ID_1'),
    (:'USER_HEAT_MANAGEMENT_ID_2', :'HEAT_MANAGEMENT_ID_2', :'USER_ID_1');

-- user_bottles
\set USER_BOTTLE_ID_1 9baf6a6c-87d6-49d6-8866-761ec4f9d8ce
\set USER_BOTTLE_ID_2 d063e36d-4b38-4003-b548-a1bd861ebf16

INSERT INTO
    user_bottles (id, bottle_id, user_id)
VALUES
    (:'USER_BOTTLE_ID_1', :'BOTTLE_ID_1', :'USER_ID_1'),
    (:'USER_BOTTLE_ID_2', :'BOTTLE_ID_2', :'USER_ID_1');
