package models

import (
	"time"
)

// Lead holds the lead's data
type Lead struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CompanyName string    `json:"company_name"`
	Email       string    `json:"email"`
	Phone       int       `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
}
