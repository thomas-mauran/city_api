
-- Création de la table city
CREATE TABLE IF NOT EXISTS city (
  id SERIAL PRIMARY KEY,
  department_code VARCHAR(255) NOT NULL,
  insee_code VARCHAR(255),
  zip_code VARCHAR(255),
  name VARCHAR(255) NOT NULL,
  lat FLOAT NOT NULL,
  lon FLOAT NOT NULL
);

-- Insertion des enregistrements initiaux
INSERT INTO city (department_code, insee_code, zip_code, name, lat, lon)
VALUES
  ('01', '01001', '01000', 'Bourg-en-Bresse', 46.2051, 5.2250),
  ('02', '02001', '02000', 'Laon', 49.5610, 3.6100),
  ('03', '03001', '03000', 'Moulins', 46.5669, 3.3375);

