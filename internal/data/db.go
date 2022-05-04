package data

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"

	"os"
)

type Config interface {
	GetUrl() string
	GetUser() string
}

func Db() (*sqlx.DB, error) {

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")

	port := os.Getenv("POSTGRES_PORT")
	database := os.Getenv("POSTGRES_DB")

	// NewTsPostgres creates new instance of postgres object

	const connStr = "user=%s password=%s  host=%s port=%s sslmode=disable"
	ds := fmt.Sprintf(connStr,
		user,
		pass,
		host,
		port)

	db, err := sqlx.Open("postgres", ds)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open connection to postgres")
	}

	err = migratePg(db, database)
	if err != nil {
		log.Fatal("problem with migration: ", err)
	}

	return db, nil

}

func migratePg(db *sqlx.DB, dbname string) error {
	path := os.Getenv("POSTGRES_MIGRATION_PATH")
	schema := os.Getenv("POSTGRES_MIGRATION_SCHEMA")
	table := os.Getenv("POSTGRES_MIGRATION_TABLE")
	goPath := os.Getenv("GOPATH")
	path = fmt.Sprintf("%s/src/github.com/ezuhl/eth/%s", goPath, path)
	migrationConfig := &postgres.Config{MigrationsTable: table, DatabaseName: dbname, SchemaName: schema}
	err := MigratePgSchema(db, path, migrationConfig)
	// check result
	switch err {
	case migrate.ErrNoChange:
		log.Println(nil, "pg schema is up-to-date")
	case nil:
		log.Println(nil, "pg schema was updated successfully")
	default:
		return err
	}

	return nil
}
