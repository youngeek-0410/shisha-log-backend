CREATE TABLE IF NOT EXISTS diaries (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    diary_equipments_id VARCHAR(36) NOT NULL,
    serve_text VARCHAR(255),
    sucking_text VARCHAR(255),
    temperature NUMERIC,
    humidity NUMERIC,
    creator_evaluation NUMERIC NOT NULL,
    taste_evaluation NUMERIC NOT NULL,
    creator_good_points VARCHAR(255),
    creator_bad_points VARCHAR(255),
    taste_comments VARCHAR(255),
    create_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS diary_flavors (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    user_flavor_id VARCHAR(36) NOT NULL,
    diary_id VARCHAR(36) NOT NULL,
    amount NUMERIC NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS diary_equipments (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    user_bottle_id VARCHAR(36) NOT NULL,
    user_bowl_id VARCHAR(36) NOT NULL,
    user_heat_management_id VARCHAR(36) NOT NULL,
    user_charcoal_id VARCHAR(36) NOT NULL,
    diary_image_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS diary_images (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    path VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_diaries (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    user_id VARCHAR(36) NOT NULL,
    diary_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS flavors (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    brand_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    create_area VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS flavor_brands (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS bottles (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    brand_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS bottle_brands (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS bowls (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    brand_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS bowl_brands (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS heat_managements (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    brand_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS heat_management_brands (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS charcoals (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    brand_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS charcoal_brands (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_flavors (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    flavor_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, flavor_id, user_id)
);

CREATE TABLE IF NOT EXISTS user_bottles (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    bottle_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_bowls (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    bowl_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_heat_managements (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    heat_management_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_charcoals (
    id VARCHAR(36) NOT NULL DEFAULT gen_random_uuid(),
    charcoal_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE flavors ADD FOREIGN KEY (brand_id) REFERENCES flavor_brands (id);

ALTER TABLE bottles ADD FOREIGN KEY (brand_id) REFERENCES bottle_brands (id);

ALTER TABLE bowls ADD FOREIGN KEY (brand_id) REFERENCES bowl_brands (id);

ALTER TABLE heat_managements ADD FOREIGN KEY (brand_id) REFERENCES heat_management_brands (id);

ALTER TABLE charcoals ADD FOREIGN KEY (brand_id) REFERENCES charcoal_brands (id);

ALTER TABLE user_flavors ADD FOREIGN KEY (flavor_id) REFERENCES flavors (id);
ALTER TABLE user_flavors ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_bottles ADD FOREIGN KEY (bottle_id) REFERENCES bottles (id);
ALTER TABLE user_bottles ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_bowls ADD FOREIGN KEY (bowl_id) REFERENCES bowls (id);
ALTER TABLE user_bowls ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_heat_managements ADD FOREIGN KEY (heat_management_id) REFERENCES heat_managements (id);
ALTER TABLE user_heat_managements ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_charcoals ADD FOREIGN KEY (charcoal_id) REFERENCES charcoals (id);
ALTER TABLE user_charcoals ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_diaries ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_diaries ADD FOREIGN KEY (diary_id) REFERENCES diaries (id);

ALTER TABLE diaries ADD FOREIGN KEY (diary_equipments_id) REFERENCES diary_equipments (id);

ALTER TABLE diary_flavors ADD FOREIGN KEY (diary_id) REFERENCES diaries (id);

ALTER TABLE diary_equipments ADD FOREIGN KEY (diary_image_id) REFERENCES diary_images (id);
