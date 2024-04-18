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
\set HEAT_MANAGEMENT_ID_1 ebe2596-e57c-4ee4-b4e5-e12919038b9c
\set HEAT_MANAGEMENT_ID_2 b85d2a9-2ac0-49d2-8ec2-3eeb4afce920

INSERT INTO
    heat_managements (id, brand_id, name)
VALUES
    (:'HEAT_MANAGEMENT_ID_1', :'HEAT_MANAGEMENT_BRAND_ID_1', 'ターキッシュリッド'),
    (:'HEAT_MANAGEMENT_ID_2', :'HEAT_MANAGEMENT_BRAND_ID_2', 'HMD - Black');

-- bottle_brands
\set BOTTLE_BRAND_1 f5a5109-d8ee-4f42-9eb0-5de2602cf97f
\set BOTTLE_BRAND_2 34e9e1a-0e66-4491-9c7a-7a996dbbbd58

INSERT INTO
    bottle_brands (id, name)
VALUES
    (:'BOTTLE_BRAND_1', 'Stimulation'),
    (:'BOTTLE_BRAND_2', 'Cloud one mini');

-- bottles
\set BOTTLE_ID_1 034a697-709a-42c3-8e95-371761a30cf7
\set BOTTLE_ID_2 5807a9a-c277-432e-a51e-3a465621cbbd

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
\set USER_FLAVOR_ID_1 a87a553-c463-41a3-be1a-aca18659f07f
\set USER_FLAVOR_ID_2 215cad9-ce65-41ac-87c6-f1ad22fdfd58
\set USER_FLAVOR_ID_3 5dd1be5-3d41-465a-b2c7-29724d642b14
\set USER_FLAVOR_ID_4 e1cf7e4-fc9b-4c2f-8cd9-bee5648e3546
\set USER_FLAVOR_ID_5 4bc97a3-14b1-4e51-9823-0ba0ff5003e5
\set USER_FLAVOR_ID_6 2ea146e-c11c-45b3-8801-10f4b965f9d6

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
\set USER_BOWL_ID_1 c312c85-669f-4ac0-ac8a-b4c8e8a8a000
\set USER_BOWL_ID_2 9d5f053-0f51-430f-b940-f85c8a2971e6
\set USER_BOWL_ID_3 22f3740-c440-41da-b3fd-7e55011a4147

INSERT INTO
    user_bowls (id, bowl_id, user_id)
VALUES
    (:'USER_BOWL_ID_1', :'BOWL_ID_1', :'USER_ID_1'),
    (:'USER_BOWL_ID_2', :'BOWL_ID_2', :'USER_ID_1'),
    (:'USER_BOWL_ID_3', :'BOWL_ID_3', :'USER_ID_1');

-- user_charcoals
\set USER_CHARCOAL_ID_1 74e350a-ef59-4410-adcb-4a0ece96a66e
\set USER_CHARCOAL_ID_2 262f9f0-11e4-442e-9f79-ae810b480ac5
\set USER_CHARCOAL_ID_3 c111c24-941f-4ece-b0f8-a3cf9185875d

INSERT INTO
    user_charcoals (id, charcoal_id, user_id)
VALUES
    (:'USER_CHARCOAL_ID_1', :'CHARCOAL_ID_1', :'USER_ID_1'),
    (:'USER_CHARCOAL_ID_2', :'CHARCOAL_ID_2', :'USER_ID_1'),
    (:'USER_CHARCOAL_ID_3', :'CHARCOAL_ID_3', :'USER_ID_1');

-- user_heat_managements
\set USER_HEAT_MANAGEMENT_ID_1 646e28c-faca-4d57-89d0-df31c7fb8481
\set USER_HEAT_MANAGEMENT_ID_2 7cdfb7e-8101-4d1e-ad76-e1df5679d031

INSERT INTO
    user_heat_managements (id, heat_management_id, user_id)
VALUES
    (:'USER_HEAT_MANAGEMENT_ID_1', :'HEAT_MANAGEMENT_ID_1', :'USER_ID_1'),
    (:'USER_HEAT_MANAGEMENT_ID_2', :'HEAT_MANAGEMENT_ID_2', :'USER_ID_1');

-- user_bottles
\set USER_BOTTLE_ID_1 f41e8e2-6c3f-4ae1-92d7-285012f34902
\set USER_BOTTLE_ID_2 f4578e2-54fc-45d5-a947-4731354381ec

INSERT INTO
    user_bottles (id, bottle_id, user_id)
VALUES
    (:'USER_BOTTLE_ID_1', :'BOTTLE_ID_1', :'USER_ID_1'),
    (:'USER_BOTTLE_ID_2', :'BOTTLE_ID_2', :'USER_ID_1');
