package app

import "context"

type Person struct {
	Id        int    `db:"Id"`
	Email     string `db:"Email"`
	Phone     string `db:"Phone"`
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`
}
type PersonLogic interface {
	GetList(limit int, search string, offset int, ctx context.Context) (error, []Person)
	GetPerson(GetId int, ctx context.Context) (error, []Person)
	AddPerson(Id int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []Person)
	DeletePerson(GetId int, ctx context.Context) (error, []Person)
	UpdatePerson(GetId int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []Person)
}
type PersonRepository interface {
	GetList(limit int, search string, offset int, ctx context.Context) (error, []Person)
	GetPerson(GetId int, ctx context.Context) (error, []Person)
	AddPerson(Id int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []Person)
	DeletePerson(GetId int, ctx context.Context) (error, []Person)
	UpdatePerson(GetId int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []Person)
}
