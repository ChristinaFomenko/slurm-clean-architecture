package postgres

import (
	"context"
	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/tools/transaction"
	"github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/repository/storage/postgres/dao"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/type/queryParameter"
	"github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	var ctx = context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.createContactTx(ctx, tx, contacts...)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) createContactTx(ctx context.Context, tx pgx.Tx, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	if len(contacts) == 0 {
		return []*contact.Contact{}, nil
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"slurm", "contact"},
		dao.CreateColumnContact,
		r.toCopyFromSource(contacts...))
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *Repository) UpdateContact(ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListContact(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error) {
	panic("implement me")
}

func (r *Repository) CountContact() (uint64, error) {
	panic("implement me")
}
