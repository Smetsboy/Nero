package logic

import (
	"PeopleService/app"
	"context"
)

type PersonLogic struct {
	repository app.PersonRepository
}

func NewPersonLogic(repository app.PersonRepository) *PersonLogic {
	return &PersonLogic{repository: repository}
}
func (u *PersonLogic) GetList(limit int, search string, offset int, ctx context.Context) (error, []app.Person) {
	return u.repository.GetList(limit, search, offset, ctx)
}
func (u *PersonLogic) GetPerson(GetId int, ctx context.Context) (error, []app.Person) {
	return u.repository.GetPerson(GetId, ctx)
}
func (u *PersonLogic) AddPerson(Id int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []app.Person) {
	return u.repository.AddPerson(Id, Email, Phone, FirstName, LastName, ctx)
}
func (u *PersonLogic) DeletePerson(GetId int, ctx context.Context) (error, []app.Person) {
	return u.repository.DeletePerson(GetId, ctx)
}
func (u *PersonLogic) UpdatePerson(GetId int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []app.Person) {
	return u.repository.UpdatePerson(GetId, Email, Phone, FirstName, LastName, ctx)
}
