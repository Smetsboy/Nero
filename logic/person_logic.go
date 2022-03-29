package logic

import (
	"PeopleService/app"
	"context"
)

type Logic struct {
	repository app.Repository
}

func NewLogic(repository app.Repository) *Logic {
	return &Logic{repository: repository}
}
func (u *Logic) GetList(ctx context.Context, limit int, search string, offset int) ([]app.Person, error) {
	return u.repository.GetList(ctx, limit, search, offset)
}
func (u *Logic) Get(ctx context.Context, GetId int) (app.Person, error) {
	return u.repository.Get(ctx, GetId)
}
func (u *Logic) Add(ctx context.Context, p app.Person) (app.Person, error) {
	return u.repository.Add(ctx, p)
}
func (u *Logic) Delete(ctx context.Context, GetId int) error {
	return u.repository.Delete(ctx, GetId)
}
func (u *Logic) Update(ctx context.Context, p app.Person) (app.Person, error) {
	return u.repository.Update(ctx, p)
}
