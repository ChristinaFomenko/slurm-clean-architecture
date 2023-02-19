package contact

import (
	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/type/context"
	"github.com/google/uuid"
	"time"

	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/type/queryParameter"
	"github.com/ChristinaFomenko/slurm-clean-architecture/services/contact/internal/domain/contact"
)

func (uc *UseCase) Create(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return uc.adapterStorage.CreateContact(ctx, contacts...)
}

func (uc *UseCase) Update(ctx context.Context, contactUpdate contact.Contact) (*contact.Contact, error) {
	return uc.adapterStorage.UpdateContact(ctx, contactUpdate.ID(), func(oldContact *contact.Contact) (*contact.Contact, error) {
		return contact.NewWithID(
			oldContact.ID(),
			oldContact.CreatedAt(),
			time.Now().UTC(),
			contactUpdate.PhoneNumber(),
			contactUpdate.Email(),
			contactUpdate.Name(),
			contactUpdate.Surname(),
			contactUpdate.Patronymic(),
			contactUpdate.Age(),
			contactUpdate.Gender(),
		)
	})
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteContact(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "List")
	//defer span.Finish()
	return uc.adapterStorage.ListContact(ctx, parameter)
}

func (uc *UseCase) ReadByID(ctx context.Context, ID uuid.UUID) (response *contact.Contact, err error) {
	return uc.adapterStorage.ReadContactByID(ctx, ID)
}

func (uc *UseCase) Count(ctx context.Context) (uint64, error) {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "Count")
	//defer span.Finish()
	return uc.adapterStorage.CountContact(ctx)
}
