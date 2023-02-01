package name

import "github.com/pkg/errors"

var (
	MaxLength      = 250
	ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)
)

type Name struct {
	value string
}

func (d Name) Value() string {
	return d.value
}

func New(description string) (Name, error) {
	if len([]rune(description)) > MaxLength {
		return Name{}, ErrWrongLength
	}
	return Name{value: description}, nil
}
