package group

import (
	"github.com/google/uuid"

	"github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/domain/contact"
)

func (uc *UseCase) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	panic("implement me")
}
