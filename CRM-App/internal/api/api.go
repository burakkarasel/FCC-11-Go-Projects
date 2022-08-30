package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/burakkarasel/CRM-App/internal/db"
	"github.com/gofiber/fiber/v2"
)

// ListLeads return all leads
func ListLeads(c *fiber.Ctx) error {
	leads, err := db.ListLeads()

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).Send([]byte(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	return c.JSON(leads)
}

// GetLead handles get specific lead
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	lead, err := db.GetLead(int64(ID))

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).Send([]byte(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	return c.JSON(lead)
}

// NewLead creates a new lead in DB
func NewLead(c *fiber.Ctx) error {
	var req db.NewLeadRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	l, err := db.NewLead(req)

	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	return c.JSON(l)
}

// DeleteLead deletes a specific lead in the DB
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).Send([]byte(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	return db.DeleteLead(int64(ID))
}
