package contact

import "github.com/google/uuid"

type Contact struct {
	ID         uuid.UUID
	FirstName  string
	LastName   string
	MiddleName string
}

type Groups struct {
	ID int
}
