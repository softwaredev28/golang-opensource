// file schema/migrate.go
package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
    {
        Version:     1,
        Description: "Add cities",
        Script: `
CREATE TABLE cities (
    id   serial NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);`,
    },
}

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sql.DB) error {
    driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})

    d := darwin.New(driver, migrations, nil)

    return d.Migrate()
}