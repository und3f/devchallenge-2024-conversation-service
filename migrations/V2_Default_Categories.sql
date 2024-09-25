--liquibase formatted sql

--changeset V2_Default_Categories:insertCategories
INSERT INTO categories(title)
OVERRIDING SYSTEM VALUE
VALUES
  ('Visa and Passport Services'),
  ('Diplomatic Inquiries'),
  ('Travel Advisories'),
  ('Consular Assistance'),
  ('Trade and Economic Cooperation');

--changeset V2_Default_Categories:insertVisaCategories
-- Visa and Passport Services
-- 
--     Points:
--         "Border"
--         "International"
--         "Passport"
--         "Visa"
-- 
INSERT INTO points(text) VALUES('Border');
INSERT INTO points(text) VALUES('International');
INSERT INTO points(text) VALUES('Passport');
INSERT INTO points(text) VALUES('Visa');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Visa and Passport Services'),
    (SELECT id FROM points WHERE text = 'Border')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Visa and Passport Services'),
    (SELECT id FROM points WHERE text = 'International')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Visa and Passport Services'),
    (SELECT id FROM points WHERE text = 'Passport')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Visa and Passport Services'),
    (SELECT id FROM points WHERE text = 'Visa')
  );

--changeset V2_Default_Categories:insertDiplomaticInquiries
-- Diplomatic Inquiries
-- 
--     Points:
--         "Diplomacy"
--         "Diplomatic"
--         "International"
-- 
INSERT INTO points(text) VALUES('Diplomacy');
INSERT INTO points(text) VALUES('Diplomatic');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Diplomatic Inquiries'),
    (SELECT id FROM points WHERE text = 'Diplomacy')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Diplomatic Inquiries'),
    (SELECT id FROM points WHERE text = 'Diplomatic')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Diplomatic Inquiries'),
    (SELECT id FROM points WHERE text = 'International')
  )
  
--changeset V2_Default_Categories:travelAdisories
-- Travel Advisories
-- 
--     Points:
--         "Alerts"
--         "Restrictions"
--         "Safety"
--         "Travel"
-- 
INSERT INTO points(text) VALUES('Alerts');
INSERT INTO points(text) VALUES('Restrictions');
INSERT INTO points(text) VALUES('Safety');
INSERT INTO points(text) VALUES('Travel');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Travel Advisories'),
    (SELECT id FROM points WHERE text = 'Travel')
  )

--changeset V2_Default_Categories:consularAssistance
-- Consular Assistance
-- 
--     Points:
--         "Apostille"
--         "Consular"
--         "Internation Adoption"
--         "Legalization"
--         "Ministry of Foreign Affairs of Ukraine"
--         "Passport"
-- 
INSERT INTO points(text) VALUES('Apostille');
INSERT INTO points(text) VALUES('Consular');
INSERT INTO points(text) VALUES('Immigration to Ukraine');
INSERT INTO points(text) VALUES('International Adoption');
INSERT INTO points(text) VALUES('Legalization');
INSERT INTO points(text) VALUES('Ministry of Foreign Affairs of Ukraine');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Apostille')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Consular')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Immigration to Ukraine')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Ministry of Foreign Affairs of Ukraine')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Ministry of Foreign Affairs of Ukraine')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Legalization')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'International Adoption')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Consular Assistance'),
    (SELECT id FROM points WHERE text = 'Passport')
  )
  
--changeset V2_Default_Categories:insertTradeAndEconomicCooperation
-- Trade and Economic Cooperation
-- 
--     Points:
--         "Economic Cooperation"
--         "Economic Development"
--         "Trade"
INSERT INTO points(text) VALUES('Economic Cooperation');
INSERT INTO points(text) VALUES('Economic Development');
INSERT INTO points(text) VALUES('Trade');

INSERT INTO category_points(category_id, point_id)
VALUES
  (
    (SELECT id FROM categories WHERE title = 'Trade and Economic Cooperation' ),
    (SELECT id FROM points WHERE text = 'Economic Cooperation')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Trade and Economic Cooperation' ),
    (SELECT id FROM points WHERE text = 'Economic Development')
  ),
  (
    (SELECT id FROM categories WHERE title = 'Trade and Economic Cooperation' ),
    (SELECT id FROM points WHERE text = 'Trade')
  )
  
