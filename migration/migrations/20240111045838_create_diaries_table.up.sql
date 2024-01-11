CREATE TABLE diaries (
    id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID(), 1)),
    user_bottle_id BINARY(16) NOT NULL,
    user_bowl_id BINARY(16) NOT NULL,
    user_heat_heat_management_id BINARY(16) NOT NULL,
    user_chacoal_id BINARY(16) NOT NULL,
    process_evaluation FLOAT NOT NULL,
    creator_good_point VARCHAR(256),
    creator_bad_point VARCHAR(256),
    taste_evaluation FLOAT NOT NULL,
    taste_comment VARCHAR(256),
    create_date DATE NOT NULL,
	created_at DATETIME,
	updated_at DATETIME,
	deleted_at DATETIME,
    PRIMARY KEY(id)
);