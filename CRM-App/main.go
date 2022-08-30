package main

import (
	"database/sql"
	"log"

	"github.com/burakkarasel/CRM-App/internal/api"
	"github.com/burakkarasel/CRM-App/internal/db"
	"github.com/burakkarasel/CRM-App/internal/dsn"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	initDatabase()
	defer db.DBConn.Close()
	runDBMigration("file://internal/db/migrations", dsn.DBSource)
	// setup routes
	setupRoutes(app)
	err := app.Listen(":3000")

	if err != nil {
		log.Fatal("cannot start server at port 3000:", err)
	}
}

// initDatabase creates the db connection
func initDatabase() {
	var err error
	db.DBConn, err = sql.Open(dsn.DBDriver, dsn.DBSource)

	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
}

// setupRoutes sets up routes for the app
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", api.ListLeads)
	app.Get("/api/v1/lead/:id", api.GetLead)
	app.Post("/api/v1/lead", api.NewLead)
	app.Delete("/api/v1/lead/:id", api.DeleteLead)
}

// runDBMigration runs the migrations at the start of the program
func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated succesfully")
}
