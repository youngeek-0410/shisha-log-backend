CREATE TABLE `diaries` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  diary_equipments_id BINARY(16) NOT NULL,
  serve_text VARCHAR(255),
  sucking_text VARCHAR(255),
  temperature FLOAT,
  humidity FLOAT,
  creator_evaluation FLOAT NOT NULL,
  taste_evaluation FLOAT NOT NULL,
  creator_good_points VARCHAR(255),
  creator_bad_points VARCHAR(255),
  taste_comments VARCHAR(255),
  create_date DATE NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_flavors` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  user_flavor_id BINARY(16) NOT NULL,
  diary_id BINARY(16) NOT NULL,
  amount FLOAT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_equipments` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  user_bottle_id BINARY(16) NOT NULL,
  user_bowl_id BINARY(16) NOT NULL,
  user_heat_management_id BINARY(16) NOT NULL,
  user_charcoal_id BINARY(16) NOT NULL,
  diary_image_id BINARY(16) NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `diary_images` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  path VARCHAR(255),
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_diaries` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  user_id BINARY(16) NOT NULL,
  diary_id BINARY(16) NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `flavors` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  create_area VARCHAR(255),
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `flavor_brands` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bottles` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bottle_brands`(
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bowls` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `bowl_brands` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `heat_managements` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `heat_management_brands` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `charcoals` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  brand_id BINARY(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `charcoal_brands` (
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

CREATE TABLE `user_flavors` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  flavor_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_bottles` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  bottle_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_bowls` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  bowl_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_heat_managements` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  heat_management_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE `user_charcoals` (
  id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
  charcoal_id BINARY(16) NOT NULL,
  user_id BINARY(16) NOT NULL,
  created_at DATETIME,
  PRIMARY KEY(id)
);

ALTER TABLE `flavors` ADD FOREIGN KEY (`brand_id`) REFERENCES `flavor_brands` (`id`);
ALTER TABLE `bottles` ADD FOREIGN KEY (`brand_id`) REFERENCES `bottle_brands` (`id`);
ALTER TABLE `bowls` ADD FOREIGN KEY (`brand_id`) REFERENCES `bowl_brands` (`id`);
ALTER TABLE `heat_managements` ADD FOREIGN KEY (`brand_id`) REFERENCES `heat_management_brands` (`id`);
ALTER TABLE `charcoals` ADD FOREIGN KEY (`brand_id`) REFERENCES `charcoal_brands` (`id`);
ALTER TABLE `user_flavors` ADD FOREIGN KEY (`flavor_id`) REFERENCES `flavors` (`id`);
ALTER TABLE `user_bottles` ADD FOREIGN KEY (`bottle_id`) REFERENCES `bottles` (`id`);
ALTER TABLE `user_bowls` ADD FOREIGN KEY (`bowl_id`) REFERENCES `bowls` (`id`);
ALTER TABLE `user_heat_managements` ADD FOREIGN KEY (`heat_management_id`) REFERENCES `heat_managements` (`id`);
ALTER TABLE `user_charcoals` ADD FOREIGN KEY (`charcoal_id`) REFERENCES `charcoals` (`id`);
ALTER TABLE `user_flavors` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `user_bottles` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `user_bowls` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `user_heat_managements` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `user_charcoals` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `user_diaries` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `user_diaries` ADD FOREIGN KEY (`diary_id`) REFERENCES `diaries` (`id`);
ALTER TABLE `diaries` ADD FOREIGN KEY (`diary_equipments_id`) REFERENCES `diary_equipments` (`id`);
ALTER TABLE `diary_flavors` ADD FOREIGN KEY (`user_flavor_id`) REFERENCES `user_flavors` (`flavor_id`);
ALTER TABLE `diary_flavors` ADD FOREIGN KEY (`diary_id`) REFERENCES `diaries` (`id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_bottle_id`) REFERENCES `user_bottles` (`bottle_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_bowl_id`) REFERENCES `user_bowls` (`bowl_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_heat_management_id`) REFERENCES `user_heat_managements` (`heat_management_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`user_charcoal_id`) REFERENCES `user_charcoals` (`charcoal_id`);
ALTER TABLE `diary_equipments` ADD FOREIGN KEY (`diary_image_id`) REFERENCES `diary_images` (`id`);
