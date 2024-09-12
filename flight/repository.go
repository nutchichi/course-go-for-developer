package flight

import (
	"database/sql"
)

type FlightDB struct {
	ID          int    `db:"id"`
	Number      int    `db:"airline_number"`
	AirlineCode string `db:"airline_code"`
	Destination string `db:"destination"`
	Arrival     string `db:"arrival"`
}

type flightRepositoryDB struct {
	db *sql.DB
}

type FlightRepository interface {
	GetAll() ([]FlightDB, error)
	GetById(int) (*FlightDB, error)
	Create(FlightDB) error
	UpdateById(int, FlightDB) (FlightDB, error)
	DeleteById(int) error
}

func NewFlightRepositoryDB(db *sql.DB) FlightRepository {
	return flightRepositoryDB{db: db}
}

func (r flightRepositoryDB) GetAll() ([]FlightDB, error) {

	row, err := r.db.Query("select id, airline_number , airline_code, destination, arrival from flights;")
	if err != nil {
		return nil, err
	}

	var flights []FlightDB
	for row.Next() {
		var f FlightDB
		err := row.Scan(&f.ID, &f.Number, &f.AirlineCode, &f.Destination, &f.Arrival)
		if err != nil {
			return nil, err
		}
		flights = append(flights, f)
	}

	return flights, nil
}

func (r flightRepositoryDB) GetById(id int) (*FlightDB, error) {
	flight := FlightDB{}
	query := "SELECT id, airline_number, airline_code, destination, arrival FROM flights WHERE id=$1;"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&flight.ID, &flight.Number, &flight.AirlineCode, &flight.Destination, &flight.Arrival)
	if err != nil {
		return nil, err
	}
	return &flight, nil
}

func (r flightRepositoryDB) UpdateById(id int, flight FlightDB) (FlightDB, error) {
	// var f FlightDB
	query := "UPDATE flights SET airline_number=$1, airline_code=$2, destination=$3, arrival=$4 WHERE id=$5 RETURNING id, airline_number, airline_code, destination, arrival;"

	row := r.db.QueryRow(
		query,
		flight.Number,
		flight.AirlineCode,
		flight.Destination,
		flight.Arrival,
		id)

	err := row.Scan(&flight.ID, &flight.Number, &flight.AirlineCode, &flight.Destination, &flight.Arrival)
	if err != nil {
		return FlightDB{}, err
	}

	return flight, err
}

func (r flightRepositoryDB) DeleteById(id int) error {
	query := "DELETE FROM flights WHERE id=$1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r flightRepositoryDB) Create(flight FlightDB) error {
	query := `
    INSERT INTO flights (airline_number, airline_code, destination, arrival)
    VALUES ($1, $2, $3, $4);`

	_, err := r.db.Exec(
		query,
		flight.Number,
		flight.AirlineCode,
		flight.Destination,
		flight.Arrival,
	)

	return err
}
