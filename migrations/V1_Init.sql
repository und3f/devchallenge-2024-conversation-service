--liquibase formatted sql

--changeset V1_Init:createPoints
CREATE TABLE points (
  id    BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  text  TEXT
          UNIQUE
          NOT NULL
          CONSTRAINT point_text_min_length CHECK (char_length(text) >= 3),
    
  text_tsquery   tsquery
                    GENERATED ALWAYS
                      AS (phraseto_tsquery('english', text)) STORED
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
  text_tsvector   tsvector
                    GENERATED ALWAYS
                      AS (to_tsvector('english', text)) STORED,

  name            VARCHAR(255),
  location        VARCHAR(255),
  emotional_tone  VARCHAR(20)
);

--changeset V1_Init:createCallsFullTextSearchIndex
CREATE INDEX calls_text_tsvector_idx
  ON calls USING GIN (text_tsvector);
