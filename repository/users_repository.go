package repository

import (
	"fmt"
	"gitlab.com/layunne/users-crud-go/models"
	"sync"
)

type UsersRepository interface {
	Get(id string) *models.User
	GetAll() []*models.User
	Save(user *models.User)
	Delete(id string)
}

func NewUsersRepository() UsersRepository {
	data := make(map[string]*models.User)
	return &usersRepository{data: data}
}

type usersRepository struct {
	sync.RWMutex
	data map[string]*models.User
}

func (r *usersRepository) Get(id string) *models.User {
	r.Lock()
	defer r.Unlock()
	return r.data[id]
}

func (r *usersRepository) GetAll() []*models.User {
	r.Lock()
	users := make([]*models.User, 0, len(r.data))
	for _, v := range r.data {
		users = append(users, v)
	}
	r.Unlock()
	return users
}

func (r *usersRepository) Save(user *models.User) {
	r.Lock()
	r.data[user.Id] = user
	r.Unlock()
}

func (r *usersRepository) Delete(id string) {
	r.Lock()
	delete(r.data, id)
	r.Unlock()
}

func (r *usersRepository) getKey(id string) string {

	return fmt.Sprintf("users:%v", id)
}


