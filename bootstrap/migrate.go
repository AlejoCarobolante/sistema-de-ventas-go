package bootstrap

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	migrate_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
)

const NO_CHANGE = "no change"

func autoMigrate(dsn string, env *Env) {
	db, _ := sql.Open("mysql", dsn)
	defer db.Close()
	driver, _ := migrate_mysql.WithInstance(db, &migrate_mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		env.MigrationsFolder,
		"mysql",
		driver,
	)

	//Ambos deberian ser un fatal error en prod
	if err != nil {
		log.Fatal("Error running migrations: ", err.Error())
		return
	}
	err = m.Up()
	if err != nil && err.Error() != NO_CHANGE {
		log.Fatal("Error while running migrations: ", err.Error())
	}
}
