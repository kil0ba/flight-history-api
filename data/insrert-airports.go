package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

// Airport represents the structure of the CSV and the database table
type Airport struct {
	ID        int
	Name      string
	Code      string
	City      string
	Country   string
	Latitude  float64
	Longitude float64
	Timezone  string
}

// parseCSV parses the CSV file and returns a slice of Airport structs
func parseCSV(filePath string) ([]Airport, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // allows variable number of fields per record

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var airports []Airport
	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Invalid ID at row %d: %v", i, err)
			continue
		}

		latitude, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			log.Printf("Invalid Latitude at row %d: %v", i, err)
			continue
		}

		longitude, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			log.Printf("Invalid Longitude at row %d: %v", i, err)
			continue
		}

		airport := Airport{
			ID:        id,
			Name:      record[1],
			Code:      record[4],
			City:      record[2],
			Country:   record[3],
			Latitude:  latitude,
			Longitude: longitude,
			Timezone:  record[9],
		}

		airports = append(airports, airport)
	}

	return airports, nil
}

// insertAirports inserts the parsed data into the PostgreSQL database
func insertAirports(tx pgx.Tx, airports []Airport) error {
	for _, airport := range airports {
		_, err := tx.Exec(context.Background(),
			"INSERT INTO airports (ID, name, code, city, country, latitude, longitude, timezone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT (ID) DO NOTHING",
			airport.ID, airport.Name, airport.Code, airport.City, airport.Country, airport.Latitude, airport.Longitude, airport.Timezone)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://localhost/flighthistory?sslmode=disable&user=root&password=example")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Ensure the table exists
	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS airports (
			ID INT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			code VARCHAR(10) NOT NULL,
			city VARCHAR(255) NOT NULL,
			country VARCHAR(255) NOT NULL,
			latitude DECIMAL(9,6) NOT NULL,
			longitude DECIMAL(9,6) NOT NULL,
			timezone SMALLINT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	// Parse the CSV file
	filePath := "data/airportsShort.csv"
	airports, err := parseCSV(filePath)
	if err != nil {
		log.Fatalf("Error parsing CSV: %v\n", err)
	}

	// Start a transaction
	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("Unable to start transaction: %v\n", err)
	}

	// Insert the parsed data into the database
	err = insertAirports(tx, airports)
	if err != nil {
		tx.Rollback(context.Background())
		log.Fatalf("Error inserting data into database, rolling back: %v\n", err)
	}

	// Commit the transaction
	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatalf("Unable to commit transaction: %v\n", err)
	}

	fmt.Println("Data inserted successfully!")
}
