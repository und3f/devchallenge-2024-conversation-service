--liquibase formatted sql

--changeset V3_Sample_Call:insertSampleCall
INSERT INTO calls(processed, name, location, emotional_tone, text)
OVERRIDING SYSTEM VALUE
VALUES (
  TRUE,
  'Sample Call',
  'Kyiv',
  'Neutral',
  'Hello and welcome to out call in Kyiv. I am happy to talk about visa and diplomatic inquries!'
);

--changeset V3_Sample_Call:insertSampleCall2
INSERT INTO calls(processed, name, location, emotional_tone, text)
OVERRIDING SYSTEM VALUE
VALUES (
  TRUE,
  'Stieve McZay',
  'Kyiv',
  'Positive',
  'Good day! I am Stieve McZay from Kyiv! I need to prolong my Visa. Since I am International vistor I need consular advice! Thank you!'
);

