package main

import (
	"context"
	"log"
	"time"

	"github.com/burakkarasel/HRMS-App/internal/api"
	"github.com/burakkarasel/HRMS-App/internal/models"
	"github.com/burakkarasel/HRMS-App/internal/routes"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mg models.MongoInstance

const dbName = "hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

func main() {
	err := Connect()
	api.GetDB(&mg)

	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	app := fiber.New()

	routes.SetUpRoutes(app)

	err = app.Listen(":3000")

	if err != nil {
		log.Fatal("cannot start app at port :3000", err)
	}
}

// Connect connects to DB
func Connect() error {
	// here we created a new client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		return err
	}

	// we added timeout for launch of mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// then we connect
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	// then we connect to the db for the client
	db := client.Database(dbName)

	// finally we set mg
	mg.Client = client
	mg.DB = db

	return nil
}
