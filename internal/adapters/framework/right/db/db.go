package db

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/sinakeshmiri/asset-notifier/internal/ports"

	// blank import for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Adapter implements the DbPort interface
type Adapter struct {
	db *sql.DB
}

// NewAdapter creates a new Adapter
func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

// CloseDbConnection closes the db  connection
func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

// Adds  record  to the  Database
func (da Adapter) AddRecord(rec ports.NSrecord) error {
	//DNS_RECORDS
	queryString, args, err := sq.Insert("dns_records").Columns("provider", "type", "name", "value").
		Values(rec.Provider, rec.Type, rec.Name, rec.Value).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}

// Deletes a  record  from the  Database
func (da Adapter) DeleteRecord(rec ports.NSrecord) error {
	deltedRecs := sq.Delete("*").From("dns_records").Where(
		sq.Eq{"provider": rec.Provider, "type": rec.Type, "name": rec.Name, "value": rec.Value})

	queryString, args, err := deltedRecs.ToSql()

	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}

// Return all  avaiable records
func (da Adapter) GetRecords() ([]ports.NSrecord, error) {
	var res []ports.NSrecord
	allRecs := sq.Select("*").From("dns_records")
	queryString, args, err := allRecs.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := da.db.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var nsr ports.NSrecord
		rows.Scan(&nsr.Provider, &nsr.Type, &nsr.Name, &nsr.Value)
		res = append(res, nsr)
	}

	return res, nil
}
