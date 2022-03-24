package PostBD

import (
	"PeopleService/app"
	"context"
	"fmt"
	"github.com/gocraft/dbr/v2"
)

type PersonRepository struct {
}

func NewUrlRepository() *PersonRepository {
	return &PersonRepository{}
}
func (*PersonRepository) GetList(limit int, search string, offset int, ctx context.Context) (err error, person []app.Person) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.Select("*").From("Person").Load(&person)
	db.Select("*").From("Person").Where("Email = '" + fmt.Sprint(search) + "'").
		Offset(uint64(offset)).Limit(uint64(limit)).LoadOne(&[]app.Person{})
	return err, person
}
func (*PersonRepository) GetPerson(GetId int, ctx context.Context) (err error, person []app.Person) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.Select("*").From("Person\\").Where("Id\\ = '" + fmt.Sprint(GetId) + "'").Load(&person)

	return err, person
}
func (*PersonRepository) AddPerson(Id int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (err error, person []app.Person) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.InsertInto("Person").Columns("Id").
		Record(&Id).Columns("Email\\").Record(&Email).Columns("Phone\\").Record(&Phone).Columns("FirstName\\").
		Record(FirstName).Columns("LastName\\").Record(LastName).
		Exec()
	return err, person
}
func (*PersonRepository) DeletePerson(GetId int, ctx context.Context) (err error, person []app.Person) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.DeleteFrom("Person").Where("Id\\ = '" + fmt.Sprint(GetId) + "'").Exec()
	return err, person
}
func (*PersonRepository) UpdatePerson(GetId int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (err error, person []app.Person) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.Update("Person\\").Set("Email\\", Email).Set("Phone\\", Phone).Set("FirstName\\", FirstName).
		Set("#{\"LastName\"}", LastName).Exec()
	return err, person
}
