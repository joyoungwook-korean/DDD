package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/domain/entity"
	"github.com/percybolmer/ddd-go/domain/vo"
)

var (
	ErrInvalidPerson = errors.New("Cannot Create Person")
)

type Customer struct {
	Person      *entity.Person
	products    []*entity.Item
	Transaction []vo.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		Person:      person,
		products:    make([]*entity.Item, 0),
		Transaction: make([]vo.Transaction, 0),
	}, nil

}
