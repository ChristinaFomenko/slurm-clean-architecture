package group

import (
	"github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/domain/group/description"
	"github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/domain/group/name"
	"github.com/google/uuid"
	"time"
)

type Group struct {
	id           uuid.UUID
	createdAt    time.Time
	modifiedAt   time.Time
	name         name.Name
	description  description.Description
	contactCount uint64
}

func NewWithID(id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	name name.Name,
	description description.Description,
	contactCount uint64) *Group {
	return &Group{
		id:           id,
		createdAt:    createdAt,
		modifiedAt:   modifiedAt,
		name:         name,
		description:  description,
		contactCount: contactCount,
	}
}

func New(name name.Name, description description.Description) *Group {
	var timeNow = time.Now().UTC()
	return &Group{
		id:          uuid.New(),
		name:        name,
		description: description,
		createdAt:   timeNow,
		modifiedAt:  timeNow,
	}
}

func (g Group) ContactCount() uint64 {
	return g.contactCount
}

func (g Group) ID() uuid.UUID {
	return g.id
}

func (g Group) CreatedAt() time.Time {
	return g.createdAt
}

func (g Group) ModifiedAt() time.Time {
	return g.modifiedAt
}

func (g Group) Name() name.Name {
	return g.name
}

func (g Group) Description() description.Description {
	return g.description
}