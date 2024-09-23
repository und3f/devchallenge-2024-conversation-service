--liquibase formatted sql

--changeset V2_Default_Categories:createCategories
INSERT INTO categories(title)
OVERRIDING SYSTEM VALUE
VALUES
  ('Visa and Passport Services'),
  ('Diplomatic Inquiries'),
  ('Travel Advisories'),
  ('Consular Assistance'),
  ('Trade and Economic Cooperation');

--changeset V2_Default_Categories:insertVisaCategories
INSERT INTO points(text) VALUES('Visa');
INSERT INTO points(text) VALUES('Passport Service');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Visa and Passport Services'),
    (SELECT id FROM points WHERE text = 'Visa')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Visa and Passport Services'),
    (SELECT id FROM points WHERE text = 'Passport Service')
  );
