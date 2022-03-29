package app

import "context"

type Person struct {
	Id        int    `db:"id"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}
type Logic interface {
	GetList(context.Context, int, string, int) ([]Person, error)
	Get(context.Context, int) (Person, error)
	Add(context.Context, Person) (Person, error)
	Delete(context.Context, int) error
	Update(context.Context, Person) (Person, error)
}
type Repository interface {
	GetList(context.Context, int, string, int) ([]Person, error)
	Get(context.Context, int) (Person, error)
	Add(context.Context, Person) (Person, error)
	Delete(context.Context, int) error
	Update(context.Context, Person) (Person, error)
}
