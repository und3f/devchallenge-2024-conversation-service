--liquibase formatted sql

--changeset V1_Init:createPoints
CREATE TABLE points (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  text TEXT
);

--changeset V1_Init:createCategories
CREATE TABLE categories (
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title TEXT
);

--changeset V1_Init:createCategoryPoints
CREATE TABLE category_points (
  category_id INT REFERENCES points(id),
  point_id INT REFERENCES categories(id)
);

--changeset V1_Init:createCalls
CREATE TABLE calls (
  id              INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name            VARCHAR(255),
  location        VARCHAR(255),
  emotional_tone  VARCHAR(20),
  text            TEXT
);

--changeset V1_Init:createCallCategories
CREATE TABLE call_categories (
  call_id     INT REFERENCES calls(id),
  catogory_id INT REFERENCES categories(id)
);

