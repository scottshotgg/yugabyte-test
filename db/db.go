package db

import (
	"github.com/google/uuid"
)

type (
	DB interface {
		NewEmployee(e *Employee) (uuid.UUID, error)
		GetEmployeeByID(id uuid.UUID) (*Employee, error)
		Close() error
	}

	Employee struct {
		ID       uuid.UUID `json:",omitempty"`
		Name     string    `json:",omitempty"`
		Age      int       `json:",omitempty"`
		Language string    `json:",omitempty"`
	}
)
