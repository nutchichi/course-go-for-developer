package flight

import (
	"database/sql"
)

type FtDB struct {
	ID          int    `db:"id"`
	Number      int    `db:"airline_number"`
	AirlineCode string `db:"airline_code"`
	Destination string `db:"destination"`
	Arrival     string `db:"arrival"`
}

type ftRepository struct {
	db *sql.DB
}

type FtRepository interface {
	GetAll() ([]FtDB, error)
	GetById(int) (*FtDB, error)
	Create(FtDB) error
	UpdateById(int, FtDB) (FtDB, error)
	DeleteById(int) error
}

func NewFlightRepositoryDB(db *sql.DB) FtRepository {
	return ftRepository{db: db}
}

func (r ftRepository) GetAll() ([]FtDB, error) {
	query := "select id, airline_number , airline_code, destination, arrival from flights;"
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var ftDBs []FtDB
	for row.Next() {
		var ftDB FtDB
		err := row.Scan(&ftDB.ID, &ftDB.Number, &ftDB.AirlineCode, &ftDB.Destination, &ftDB.Arrival)
		if err != nil {
			return nil, err
		}
		ftDBs = append(ftDBs, ftDB)
	}

	return ftDBs, nil
}

func (r ftRepository) GetById(id int) (*FtDB, error) {
	ftDB := FtDB{}
	query := "SELECT id, airline_number, airline_code, destination, arrival FROM flights WHERE id=$1;"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&ftDB.ID, &ftDB.Number, &ftDB.AirlineCode, &ftDB.Destination, &ftDB.Arrival)
	if err != nil {
		return nil, err
	}
	return &ftDB, nil
}

func (r ftRepository) UpdateById(id int, ftDB FtDB) (FtDB, error) {
	query := "UPDATE flights SET airline_number=$1, airline_code=$2, destination=$3, arrival=$4 WHERE id=$5 RETURNING id, airline_number, airline_code, destination, arrival;"
	row := r.db.QueryRow(
		query,
		ftDB.Number,
		ftDB.AirlineCode,
		ftDB.Destination,
		ftDB.Arrival,
		id)

	err := row.Scan(&ftDB.ID, &ftDB.Number, &ftDB.AirlineCode, &ftDB.Destination, &ftDB.Arrival)
	if err != nil {
		return FtDB{}, err
	}

	return ftDB, err
}

func (r ftRepository) DeleteById(id int) error {
	query := "DELETE FROM flights WHERE id=$1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r ftRepository) Create(ftDB FtDB) error {
	query := `
    INSERT INTO flights (airline_number, airline_code, destination, arrival)
    VALUES ($1, $2, $3, $4);`

	_, err := r.db.Exec(
		query,
		ftDB.Number,
		ftDB.AirlineCode,
		ftDB.Destination,
		ftDB.Arrival,
	)

	return err
}
