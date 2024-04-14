CREATE TABLE IF NOT EXISTS planes (
    ID INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    IATA_code VARCHAR(5),
    ICAO_code VARCHAR(5),
    manufacturer VARCHAR(255),
    country VARCHAR(255)
);