--liquibase formatted sql

--changeset V1_Init:createPoints
CREATE TABLE points (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  text TEXT
    UNIQUE
    NOT NULL
    CONSTRAINT point_text_min_length CHECK (char_length(text) >= 3)
);

--changeset V1_Init:createCategories
CREATE TABLE categories (
  id BIGINT
      GENERATED ALWAYS
        AS IDENTITY (MINVALUE 1001)
      PRIMARY KEY,
  title TEXT
    NOT NULL
    UNIQUE
    CONSTRAINT category_title_min_length CHECK (char_length(title) >= 3)
);

--changeset V1_Init:createCategoryPoints
CREATE TABLE category_points (
  category_id BIGINT
    REFERENCES categories(id)
    ON DELETE CASCADE,
  point_id BIGINT
    REFERENCES points(id)
    ON DELETE CASCADE
);

--changeset V1_Init:createCalls
CREATE TABLE calls (
  id              BIGINT
                    GENERATED ALWAYS
                      AS IDENTITY (MINVALUE 1001)
                    PRIMARY KEY,

  processed       BOOLEAN DEFAULT FALSE,
  error           TEXT,
  text            TEXT,

  name            VARCHAR(255),
  location        VARCHAR(255),
  emotional_tone  VARCHAR(20)
);

--changeset V1_Init:createCallCategories
CREATE TABLE call_categories (
  call_id     BIGINT
    REFERENCES calls(id)
    ON DELETE CASCADE,
  category_id BIGINT
    REFERENCES categories(id)
    ON DELETE CASCADE
);

