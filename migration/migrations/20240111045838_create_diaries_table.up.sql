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

CREATE TABLE `diary_fravors` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  user_fravor_id BINARY(16) NOT NULL,
  diary_id BINARY(16) NOT NULL,
  amount SMALLINT UNSIGNED NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_equipments` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  diary_fravors_id BINARY(16) NOT NULL,
  user_bottle_id BINARY(16) NOT NULL,
  user_bowl_id BINARY(16) NOT NULL,
  user_heat_management_id BINARY(16) NOT NULL,
  user_charcoal_id BINARY(16) NOT NULL,
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

CREATE TABLE `fravor` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  create_area VARCHAR(255),
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `fravor_brand` (
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

CREATE TABLE `users` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_fravor` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  fravor_id BINARY(16) NOT NULL,
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

ALTER TABLE `fravor` ADD FOREIGN KEY (`brand_id`) REFERENCES `fravor_brand` (`id`);

ALTER TABLE `bottle` ADD FOREIGN KEY (`brand_id`) REFERENCES `bottle_brand` (`id`);

ALTER TABLE `bowl` ADD FOREIGN KEY (`brand_id`) REFERENCES `bowl_brand` (`id`);

ALTER TABLE `heat_management` ADD FOREIGN KEY (`brand_id`) REFERENCES `heat_management_brand` (`id`);

ALTER TABLE `charcoal` ADD FOREIGN KEY (`brand_id`) REFERENCES `charcoal_brand` (`id`);

ALTER TABLE `user_fravor` ADD FOREIGN KEY (`fravor_id`) REFERENCES `fravor` (`id`);

ALTER TABLE `user_bottle` ADD FOREIGN KEY (`bottle_id`) REFERENCES `bottle` (`id`);

ALTER TABLE `user_bowl` ADD FOREIGN KEY (`bowl_id`) REFERENCES `bowl` (`id`);

ALTER TABLE `user_heat_management` ADD FOREIGN KEY (`heat_management_id`) REFERENCES `heat_management` (`id`);

ALTER TABLE `user_charcoal` ADD FOREIGN KEY (`charcoal_id`) REFERENCES `charcoal` (`id`);

ALTER TABLE `user_fravor` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_bottle` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_bowl` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_heat_management` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_charcoal` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `diary_list` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `diary_list` ADD FOREIGN KEY (`diary_id`) REFERENCES `diary` (`id`);

ALTER TABLE `diary` ADD FOREIGN KEY (`diary_equipments_id`) REFERENCES `diary_equipments` (`id`);

ALTER TABLE `diary_fravors` ADD FOREIGN KEY (`user_fravor_id`) REFERENCES `user_fravor` (`fravor_id`);

ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`diary_fravors_id`) REFERENCES `diary_fravors` (`id`);

ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_bottle_id`) REFERENCES `user_bottle` (`bottle_id`);

ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_bowl_id`) REFERENCES `user_bowl` (`bowl_id`);

ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_heat_management_id`) REFERENCES `user_heat_management` (`heat_management_id`);

ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_charcoal_id`) REFERENCES `user_charcoal` (`charcoal_id`);
