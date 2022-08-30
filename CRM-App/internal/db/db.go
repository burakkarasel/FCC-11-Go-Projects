package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/burakkarasel/CRM-App/internal/models"
	_ "github.com/lib/pq"
)

var (
	DBConn *sql.DB
)

// GetLead gets a lead for given id from the DB
func GetLead(id int64) (models.Lead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	query := `SELECT id, name, company_name, email, phone, created_at
		FROM leads
		WHERE id=$1
	`

	var lead models.Lead

	row := DBConn.QueryRowContext(ctx, query, id)

	err := row.Scan(&lead.ID, &lead.Name, &lead.CompanyName, &lead.Email, &lead.Phone, &lead.CreatedAt)

	if err != nil {
		return lead, err
	}

	return lead, nil
}

// ListLeads returns the all leads from the DB
func ListLeads() ([]models.Lead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	query := `SELECT id, name, company_name, email, phone, created_at
		FROM leads
	`

	var leads = []models.Lead{}

	rows, err := DBConn.QueryContext(ctx, query)

	if err != nil {
		return leads, err
	}

	for rows.Next() {
		var lead models.Lead
		if err := rows.Scan(&lead.ID, &lead.Name, &lead.CompanyName, &lead.Email, &lead.Phone, &lead.CreatedAt); err != nil {
			return leads, err
		}
		leads = append(leads, lead)
	}

	if err := rows.Err(); err != nil {
		return leads, err
	}

	return leads, nil
}

// NewLeadRequest holds the data to create a new lead
type NewLeadRequest struct {
	Name        string `json:"name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Phone       int    `json:"phone"`
}

// NewLead creates a new lead in the DB and returns it
func NewLead(arg NewLeadRequest) (models.Lead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	query := `INSERT INTO leads(name, company_name, email, phone, created_at)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, name, company_name, email, phone, created_at
	`

	var lead models.Lead

	row := DBConn.QueryRowContext(ctx, query, arg.Name, arg.CompanyName, arg.Email, arg.Phone, time.Now())

	err := row.Scan(&lead.ID, &lead.Name, &lead.CompanyName, &lead.Email, &lead.Phone, &lead.CreatedAt)

	if err != nil {
		return lead, err
	}

	return lead, nil
}

// DeleteLead deletes a lead from the DB for given id
func DeleteLead(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	query := `DELETE FROM leads
	WHERE id = $1
	`
	_, err := DBConn.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
