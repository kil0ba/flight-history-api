CREATE TABLE IF NOT EXISTS airlines (
    id SERIAL PRIMARY KEY, -- Unique OpenFlights identifier for this airlines
    name VARCHAR(255) NOT NULL, -- Name of the airlines
    alias VARCHAR(255), -- Alias of the airlines
    iata CHAR(2), -- 2-letter IATA code, if available
    icao CHAR(3), -- 3-letter ICAO code, if available
    callsign VARCHAR(255), -- Airline callsign
    country VARCHAR(255), -- Country or territory where the airport is located
    active boolean NOT NULL -- "Y" if the airlines is or has until recently been operational, "N" if it is defunct
);