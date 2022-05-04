package data

// nolint: golint // suppress linter for _
import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"strings"
)

// MigratePgSchema migrate postgres schema
func MigratePgSchema(db *sqlx.DB, sourceURL string, config *postgres.Config) error {

	driver, err := postgres.WithInstance(db.DB, config)
	if err != nil {
		return errors.Wrap(err, "unable to instantiate driver for pg migration")
	}

	// currently we support `file` schema
	if !strings.HasPrefix(sourceURL, "file://") {
		sourceURL = "file://" + sourceURL
	}

	migrator, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "unable to create pg migrator")
	}
	return migrator.Up()
}
