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
func (*PersonRepository) GetList(ctx context.Context, limit int, search string, offset int) (person []app.Person, err error) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	smtp := db.Select("*").From("Person")
	//db.Select("*").From("Person").Where("Email", fmt.Sprint(search)).
	//	Offset(uint64(offset)).Limit(uint64(limit)).Load(&person)
	if search != "" {
		smtp.Where("Email =" + "'" + fmt.Sprint(search) + "'")
	}
	if limit != 0 {
		smtp.Limit(uint64(limit))
	}
	if offset != 0 {
		smtp.Offset(uint64(offset))
	}
	smtp.Load(&person)
	return person, err
}
func (*PersonRepository) Get(ctx context.Context, GetId int) (person app.Person, err error) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.Select("*").From("Person").Where("Id = '" + fmt.Sprint(GetId) + "'").Load(&person)

	return person, err
}
func (*PersonRepository) Add(ctx context.Context, p app.Person) (person app.Person, err error) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.InsertInto("Person").Columns("Id").
		Record(&p.Id).Columns("Email").Record(&p.Email).Columns("Phone").Record(&p.Phone).Columns("FirstName").
		Record(&p.FirstName).Columns("LastName").Record(&p.LastName).
		Exec()
	return person, err
}
func (*PersonRepository) Delete(ctx context.Context, GetId int) (err error) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.DeleteFrom("Person").Where("Id = '" + fmt.Sprint(GetId) + "'").Exec()
	return err
}
func (*PersonRepository) Update(ctx context.Context, p app.Person) (person app.Person, err error) {
	db := ctx.Value("Postgres-client").(*dbr.Session)
	db.Update("Person").Set("Email", p.Email).Set("Phone", p.Phone).Set("FirstName", p.FirstName).
		Set("LastName", p.LastName).Exec()
	return person, err
}
