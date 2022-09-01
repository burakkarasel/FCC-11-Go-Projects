package api

import (
	"net/http"

	"github.com/burakkarasel/HRMS-App/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mg *models.MongoInstance

// GetDB gets the db in api package
func GetDB(db *models.MongoInstance) {
	mg = db
}

// ListEmployees returns all the employees
func ListEmployees(c *fiber.Ctx) error {
	var employees []models.Employee
	query := bson.D{{}}

	cursor, err := mg.DB.Collection("employees").Find(c.Context(), query)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	err = cursor.All(c.Context(), &employees)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).JSON(employees)
}

// GetEmployee returns the employee for given ID
func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	// first i get the id from URI and change it to hex
	employeeID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// then i create a read query
	filter := bson.D{{Key: "_id", Value: employeeID}}

	rec := mg.DB.Collection("employees").FindOne(c.Context(), filter)

	// if any error is occured during finding i check for error
	if rec.Err() != nil {
		if rec.Err() == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// then i decode the found record and return it
	foundEmployee := &models.Employee{}
	rec.Decode(foundEmployee)

	return c.Status(http.StatusOK).JSON(foundEmployee)
}

// NewEmployee creates a new employee in DB
func NewEmployee(c *fiber.Ctx) error {
	collection := mg.DB.Collection("employees")

	var employee models.Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	employee.ID = ""

	// first i insert the new employee to DB
	r, err := collection.InsertOne(c.Context(), employee)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// then i get the inserted employee from DB
	filter := bson.D{{Key: "_id", Value: r.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// then i decode the inserted value to JSON
	createdEmployee := &models.Employee{}
	createdRecord.Decode(createdEmployee)

	return c.Status(http.StatusCreated).JSON(createdEmployee)
}

// UpdateEmployee updates a given Employee
func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	// we convert the id given in URI to primitive
	employeeID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	var employee models.Employee

	// we parse the body to update
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// then we write 2 queriest to first find then update the record
	// find query
	query := bson.D{{Key: "_id", Value: employeeID}}
	// update query
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err = mg.DB.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

	// we check for error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// if no error occurs we return parsed body and we added the id
	employee.ID = id

	return c.Status(http.StatusOK).JSON(employee)
}

// DeleteEmployee deletes an employee for given id
func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}

	r, err := mg.DB.Collection("employees").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if r.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).SendString("record not found")
	}

	return c.Status(http.StatusOK).SendString("record deleted")
}
