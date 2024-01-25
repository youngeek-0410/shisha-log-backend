CREATE TABLE `diary` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  diary_equipments_id BINARY(16) NOT NULL,
  image_id BINARY(16),
  sucking_text VARCHAR(255),
  creator_evaluation FLOAT NOT NULL,
  taste_evaluation FLOAT NOT NULL,
  creator_good_points VARCHAR(255),
  creator_bad_points VARCHAR(255),
  taste_comments VARCHAR(255),
  is_created BOOLEAN NOT NULL,
  create_date DATE NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_flavors` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  user_flavor_id BINARY(16) NOT NULL,
  diary_id BINARY(16) NOT NULL,
  amount SMALLINT UNSIGNED NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_equipments` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  diary_flavors_id BINARY(16) NOT NULL,
  user_bottle_id BINARY(16) NOT NULL,
  user_bowl_id BINARY(16) NOT NULL,
  user_heat_management_id BINARY(16) NOT NULL,
  user_charcoal_id BINARY(16) NOT NULL,
  diary_image_id BINARY(16) NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_image` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  path VARCHAR(255),
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_list` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  user_id BINARY(16) NOT NULL,
  diary_id BINARY(16) NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `flavor` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  create_area VARCHAR(255),
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `flavor_brand` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bottle` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bottle_brand`(
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bowl` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bowl_brand` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `heat_management` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `heat_management_brand` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `charcoal` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `charcoal_brand` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_flavor` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  flavor_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_bottle` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  bottle_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_bowl` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  bowl_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_heat_management` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  heat_management_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_charcoal` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  charcoal_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

ALTER TABLE `flavor` ADD FOREIGN KEY (`brand_id`) REFERENCES `flavor_brand` (`id`);
ALTER TABLE `bottle` ADD FOREIGN KEY (`brand_id`) REFERENCES `bottle_brand` (`id`);
ALTER TABLE `bowl` ADD FOREIGN KEY (`brand_id`) REFERENCES `bowl_brand` (`id`);
ALTER TABLE `heat_management` ADD FOREIGN KEY (`brand_id`) REFERENCES `heat_management_brand` (`id`);
ALTER TABLE `charcoal` ADD FOREIGN KEY (`brand_id`) REFERENCES `charcoal_brand` (`id`);
ALTER TABLE `user_flavor` ADD FOREIGN KEY (`flavor_id`) REFERENCES `flavor` (`id`);
ALTER TABLE `user_bottle` ADD FOREIGN KEY (`bottle_id`) REFERENCES `bottle` (`id`);
ALTER TABLE `user_bowl` ADD FOREIGN KEY (`bowl_id`) REFERENCES `bowl` (`id`);
ALTER TABLE `user_heat_management` ADD FOREIGN KEY (`heat_management_id`) REFERENCES `heat_management` (`id`);
ALTER TABLE `user_charcoal` ADD FOREIGN KEY (`charcoal_id`) REFERENCES `charcoal` (`id`);
ALTER TABLE `user_flavor` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_bottle` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_bowl` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_heat_management` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_charcoal` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `diary_list` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `diary_list` ADD FOREIGN KEY (`diary_id`) REFERENCES `diary` (`id`);
ALTER TABLE `diary` ADD FOREIGN KEY (`diary_equipments_id`) REFERENCES `diary_equipments` (`id`);
ALTER TABLE `diary` ADD FOREIGN KEY (`image_id`) REFERENCES `diary_image` (`id`);
ALTER TABLE `diary_flavors` ADD FOREIGN KEY (`user_flavor_id`) REFERENCES `user_flavor` (`flavor_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`diary_flavors_id`) REFERENCES `diary_flavors` (`id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_bottle_id`) REFERENCES `user_bottle` (`bottle_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_bowl_id`) REFERENCES `user_bowl` (`bowl_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_heat_management_id`) REFERENCES `user_heat_management` (`heat_management_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_charcoal_id`) REFERENCES `user_charcoal` (`charcoal_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`diary_image_id`) REFERENCES `diary_image` (`id`);
