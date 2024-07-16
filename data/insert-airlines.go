package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

type Airline struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	IATA     string `json:"iata" validate:"max=2"`
	ICAO     string `json:"icao" validate:"max=3"`
	Callsign string `json:"callsign"`
	Country  string `json:"country"`
	Active   bool   `json:"active"`
}

func parseCSV(filePath string) ([]Airline, error) {
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

	var airlines []Airline
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

		active := false
		if record[7] == "Y" {
			active = true
		}

		airline := Airline{
			ID:       id,
			Name:     record[1],
			Alias:    record[2],
			IATA:     record[3],
			ICAO:     record[4],
			Callsign: record[5],
			Country:  record[6],
			Active:   active,
		}

		validate := validator.New()
		err = validate.Struct(airline)
		if err != nil {
			log.Printf("Validation error at row %d: %v", i, err)
			continue
		}

		airlines = append(airlines, airline)
	}

	return airlines, nil
}

func insertAirlines(tx pgx.Tx, airlines []Airline) error {
	for _, airline := range airlines {
		_, err := tx.Exec(context.Background(),
			"INSERT INTO airlines (id, name, alias, iata, icao, callsign, country, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT (id) DO NOTHING",
			airline.ID, airline.Name, airline.Alias, airline.IATA, airline.ICAO, airline.Callsign, airline.Country, airline.Active)
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

	// Parse the CSV file
	filePath := "data/airlines.csv"
	airlines, err := parseCSV(filePath)
	if err != nil {
		log.Fatalf("Error parsing CSV: %v\n", err)
	}

	// Start a transaction
	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("Unable to start transaction: %v\n", err)
	}

	// Insert the parsed data into the database
	err = insertAirlines(tx, airlines)
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
