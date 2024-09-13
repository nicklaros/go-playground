package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type MerchantBranch struct {
	BranchID    int
	Name        string
	Geolocation Point
}

type Point struct {
	Latitude  float64
	Longitude float64
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "example"
)

func connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		user, password, host, port, dbname)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func queryLocations(db *sql.DB) ([]MerchantBranch, error) {
	rows, err := db.Query("SELECT branch_id, name, geolocation FROM merchant_branch")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []MerchantBranch
	for rows.Next() {
		var loc MerchantBranch
		var point Point

		if err := rows.Scan(&loc.BranchID, &loc.Name, &point); err != nil {
			return nil, err
		}
		loc.Geolocation = point
		locations = append(locations, loc)
	}

	return locations, nil
}

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	locations, err := queryLocations(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locations:")
	for _, loc := range locations {
		fmt.Printf("ID: %d, Name: %s, Latitude: %f, Longitude: %f\n", loc.BranchID, loc.Name, loc.Geolocation.Latitude, loc.Geolocation.Longitude)
	}
}
