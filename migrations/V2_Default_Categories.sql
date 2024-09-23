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

--changeset V2_Default_Categories:insertDiplomaticInquiries
INSERT INTO points(text) VALUES('Diplomatic');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Diplomatic Inquiries'),
    (SELECT id FROM points WHERE text = 'Diplomatic')
  )
  
--changeset V2_Default_Categories:travelAdisories
INSERT INTO points(text) VALUES('Travel');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Travel Advisories'),
    (SELECT id FROM points WHERE text = 'Travel')
  )

--changeset V2_Default_Categories:consular
INSERT INTO points(text) VALUES('Consular');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Consular')
  )
  
--changeset V2_Default_Categories:insertTrade
INSERT INTO points(text) VALUES('Trade');
INSERT INTO points(text) VALUES('Economic Cooperation');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Trade and Economic Cooperation' ),
    (SELECT id FROM points WHERE text = 'Trade')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Trade and Economic Cooperation' ),
    (SELECT id FROM points WHERE text = 'Economic Cooperation')
  )
  
