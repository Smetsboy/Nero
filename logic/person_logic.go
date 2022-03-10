package logic

import (
	"PeopleService/app"
	"PeopleService/postgres"
	"context"
)

func GetList(limit int, search string, offset int, ctx context.Context) []app.Person {
	return PostBD.GetList(limit, search, offset, ctx)
}
func GetPerson(GetId int, ctx context.Context) []app.Person {
	return PostBD.GetPerson(GetId, ctx)
}
func AddPerson(Id int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) []app.Person {
	return PostBD.AddPerson(Id, Email, Phone, FirstName, LastName, ctx)
}
func DeletePerson(GetId int, ctx context.Context) []app.Person {
	return PostBD.DeletePerson(GetId, ctx)
}
func UpdatePerson(GetId int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) []app.Person {
	return PostBD.UpdatePerson(GetId, Email, Phone, FirstName, LastName, ctx)
}
