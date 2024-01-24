CREATE TABLE [flavor] (
  [id] uuid PRIMARY KEY NOT NULL,
  [brand_id] uuid NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [create_area] nvarchar(255),
  [created_at] timestamp
)
GO

CREATE TABLE [flavor_brand] (
  [id] uuid PRIMARY KEY NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [bottle] (
  [id] uuid PRIMARY KEY NOT NULL,
  [brand_id] uuid NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [bottle_brand] (
  [id] uuid PRIMARY KEY NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [bowl] (
  [id] uuid PRIMARY KEY NOT NULL,
  [brand_id] uuid NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [bowl_brand] (
  [id] uuid PRIMARY KEY NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [heat_management] (
  [id] uuid PRIMARY KEY NOT NULL,
  [brand_id] uuid NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [heat_management_brand] (
  [id] uuid PRIMARY KEY NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [charcoal] (
  [id] uuid PRIMARY KEY NOT NULL,
  [brand_id] uuid NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [charcoal_brand] (
  [id] uuid PRIMARY KEY NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [users] (
  [id] uuid PRIMARY KEY NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [user_flavor] (
  [id] uuid PRIMARY KEY NOT NULL,
  [flavor_id] uuid NOT NULL,
  [user_id] uuid NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [user_bottle] (
  [id] uuid PRIMARY KEY NOT NULL,
  [bottle_id] uuid NOT NULL,
  [user_id] uuid NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [user_bowl] (
  [id] uuid PRIMARY KEY NOT NULL,
  [bowl_id] uuid NOT NULL,
  [user_id] uuid NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [user_heat_management] (
  [id] uuid PRIMARY KEY NOT NULL,
  [heat_management_id] uuid NOT NULL,
  [user_id] uuid NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [user_charcoal] (
  [id] uuid PRIMARY KEY NOT NULL,
  [charcoal_id] uuid NOT NULL,
  [user_id] uuid NOT NULL,
  [created_at] timestamp
)
GO

CREATE TABLE [diary_list] (
  [id] uuid PRIMARY KEY NOT NULL,
  [user_id] uuid NOT NULL,
  [diary_id] uuid NOT NULL,
  [created_at] timestamp,
  [updated_at] timestamp
)
GO

CREATE TABLE [diary_equipments] (
  [id] uuid PRIMARY KEY NOT NULL,
  [diary_flavors_id] uuid NOT NULL,
  [user_bottle_id] uuid NOT NULL,
  [user_bowl_id] uuid NOT NULL,
  [user_heat_management_id] uuid NOT NULL,
  [user_charcoal_id] uuid NOT NULL,
  [diary_image_id] uuid NOT NULL,
  [created_at] timestamp,
  [updated_at] timestamp
)
GO

CREATE TABLE [diary_image] (
  [id] uuid PRIMARY KEY NOT NULL,
  [path] nvarchar(255),
  [created_at] timestamp,
  [updated_at] timestamp
)
GO

CREATE TABLE [diaries] (
  [id] uuid PRIMARY KEY NOT NULL,
  [diary_equipments_id] uuid NOT NULL,
  [sucking_text] nvarchar(255),
  [creator_evaluation] decimal NOT NULL,
  [taste_evaluation] decimal NOT NULL,
  [creator_good_points] nvarchar(255),
  [creator_bad_points] nvarchar(255),
  [taste_comments] nvarchar(255),
  [is_created] bool NOT NULL,
  [create_date] datetime NOT NULL,
  [created_at] timestamp,
  [updated_at] timestamp
)
GO

CREATE TABLE [diary_flavors] (
  [id] uuid PRIMARY KEY NOT NULL,
  [user_flavor_id] uuid NOT NULL,
  [diary_id] uuid NOT NULL,
  [amount] decimal NOT NULL,
  [created_at] timestamp,
  [updated_at] timestamp
)
GO

ALTER TABLE [flavor] ADD FOREIGN KEY ([brand_id]) REFERENCES [flavor_brand] ([id])
GO

ALTER TABLE [bottle] ADD FOREIGN KEY ([brand_id]) REFERENCES [bottle_brand] ([id])
GO

ALTER TABLE [bowl] ADD FOREIGN KEY ([brand_id]) REFERENCES [bowl_brand] ([id])
GO

ALTER TABLE [heat_management] ADD FOREIGN KEY ([brand_id]) REFERENCES [heat_management_brand] ([id])
GO

ALTER TABLE [charcoal] ADD FOREIGN KEY ([brand_id]) REFERENCES [charcoal_brand] ([id])
GO

ALTER TABLE [user_flavor] ADD FOREIGN KEY ([flavor_id]) REFERENCES [flavor] ([id])
GO

ALTER TABLE [user_bottle] ADD FOREIGN KEY ([bottle_id]) REFERENCES [bottle] ([id])
GO

ALTER TABLE [user_bowl] ADD FOREIGN KEY ([bowl_id]) REFERENCES [bowl] ([id])
GO

ALTER TABLE [user_heat_management] ADD FOREIGN KEY ([heat_management_id]) REFERENCES [heat_management] ([id])
GO

ALTER TABLE [user_charcoal] ADD FOREIGN KEY ([charcoal_id]) REFERENCES [charcoal] ([id])
GO

ALTER TABLE [user_flavor] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [user_bottle] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [user_bowl] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [user_heat_management] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [user_charcoal] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [diary_list] ADD FOREIGN KEY ([user_id]) REFERENCES [users] ([id])
GO

ALTER TABLE [diary_list] ADD FOREIGN KEY ([diary_id]) REFERENCES [diaries] ([id])
GO

ALTER TABLE [diaries] ADD FOREIGN KEY ([diary_equipments_id]) REFERENCES [diary_equipments] ([id])
GO

ALTER TABLE [diary_flavors] ADD FOREIGN KEY ([user_flavor_id]) REFERENCES [user_flavor] ([flavor_id])
GO

ALTER TABLE [diary_equipments] ADD FOREIGN KEY ([diary_flavors_id]) REFERENCES [diary_flavors] ([id])
GO

ALTER TABLE [diary_equipments] ADD FOREIGN KEY ([user_bottle_id]) REFERENCES [user_bottle] ([bottle_id])
GO

ALTER TABLE [diary_equipments] ADD FOREIGN KEY ([user_bowl_id]) REFERENCES [user_bowl] ([bowl_id])
GO

ALTER TABLE [diary_equipments] ADD FOREIGN KEY ([user_heat_management_id]) REFERENCES [user_heat_management] ([heat_management_id])
GO

ALTER TABLE [diary_equipments] ADD FOREIGN KEY ([user_charcoal_id]) REFERENCES [user_charcoal] ([charcoal_id])
GO

ALTER TABLE [diary_equipments] ADD FOREIGN KEY ([diary_image_id]) REFERENCES [diary_image] ([id])
GO
