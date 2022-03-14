package PostBD

import (
	"PeopleService/app"
	"context"
	"database/sql"
	"fmt"
)

type PersonRepository struct {
}

func NewUrlRepository() *PersonRepository {
	return &PersonRepository{}
}
func (*PersonRepository) GetList(limit int, search string, offset int, ctx context.Context) (error, []app.Person) {
	db := ctx.Value("Postgres-client").(*sql.DB)
	query := "select * from Person"
	if search != "" {
		query += " where \"Email\" = '" + fmt.Sprint(search) + "'"
	}
	if offset != 0 {
		query += " offset " + fmt.Sprint(offset)
	}
	if limit != 0 {
		query += " limit " + fmt.Sprint(limit)
	}
	rows, err := db.Query(query)
	if err != nil {
		return err, []app.Person{}
	}
	defer rows.Close()
	Persons := []app.Person{}
	for rows.Next() {
		p := app.Person{}
		err := rows.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		if err != nil {
			return err, []app.Person{}
			continue
		}
		Persons = append(Persons, p)
	}
	for _, p := range Persons {
		fmt.Println(p.Id, p.Email, p.Phone, p.FirstName, p.LastName)
	}
	defer rows.Close()
	return nil, Persons
}
func (*PersonRepository) GetPerson(GetId int, ctx context.Context) (error, []app.Person) {
	db := ctx.Value("Postgres-client").(*sql.DB)
	query := "select * from Person where \"Id\" = '" + fmt.Sprint(GetId) + "'"
	rows, err := db.Query(query)
	if err != nil {
		return err, []app.Person{}
	}
	defer rows.Close()
	Persons := []app.Person{}
	for rows.Next() {
		p := app.Person{}
		err := rows.Scan(&p.Id, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
		if err != nil {
			return err, []app.Person{}
			continue
		}
		Persons = append(Persons, p)
	}
	for _, p := range Persons {
		fmt.Println(p.Id, p.Email, p.Phone, p.FirstName, p.LastName)
	}
	defer rows.Close()
	return nil, Persons
}
func (*PersonRepository) AddPerson(Id int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []app.Person) {
	db := ctx.Value("Postgres-client").(*sql.DB)
	Persons := []app.Person{}
	result, err := db.Exec("insert into Person (\"Id\", \"Email\", \"Phone\", \"FirstName\", \"LastName\") values ($1,$2,$3,$4,$5)",
		Id, Email, Phone, FirstName, LastName)
	if err != nil {
		return err, []app.Person{}
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return nil, Persons
}
func (*PersonRepository) DeletePerson(GetId int, ctx context.Context) (error, []app.Person) {
	db := ctx.Value("Postgres-client").(*sql.DB)
	query := "delete from Person where \"Id\" = '" + fmt.Sprint(GetId) + "'"
	result, err := db.Exec(query)
	if err != nil {
		return err, []app.Person{}
	}
	Persons := []app.Person{}
	for _, p := range Persons {
		fmt.Println(p.Id, p.Email, p.Phone, p.FirstName, p.LastName)
	}
	fmt.Println(result.RowsAffected())
	return nil, Persons
}
func (*PersonRepository) UpdatePerson(GetId int, Email string, Phone string, FirstName string, LastName string, ctx context.Context) (error, []app.Person) {
	db := ctx.Value("Postgres-client").(*sql.DB)
	query := "update Person set \"Email\"= '" + fmt.Sprint(Email) + "'," + " \"Phone\"= '" + fmt.Sprint(Phone) + "'," + " \"FirstName\"= '" +
		fmt.Sprint(FirstName) + "'," + " \"LastName\"  = '" + fmt.Sprint(LastName) + "'" + " where \"Id\" = '" + fmt.Sprint(GetId) + "'"
	result, err := db.Exec(query)
	if err != nil {
		return err, []app.Person{}
	}
	Persons := []app.Person{}
	for _, p := range Persons {
		fmt.Println(p.Id, p.Email, p.Phone, p.FirstName, p.LastName)
	}
	fmt.Println(result.RowsAffected())
	return nil, Persons
}
