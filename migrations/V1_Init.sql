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
  category_id INT
    REFERENCES categories(id)
    ON DELETE CASCADE,
  point_id INT
    REFERENCES points(id)
    ON DELETE CASCADE
);

--changeset V1_Init:createCalls
CREATE TABLE calls (
  id              INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  processed       BOOLEAN DEFAULT FALSE,
  error           TEXT,
  text            TEXT,

  name            VARCHAR(255),
  location        VARCHAR(255),
  emotional_tone  VARCHAR(20)
);

--changeset V1_Init:createCallCategories
CREATE TABLE call_categories (
  call_id     INT
    REFERENCES calls(id)
    ON DELETE CASCADE,
  category_id INT
    REFERENCES categories(id)
    ON DELETE CASCADE
);

