package services

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/layunne/users-crud-go/errors"
	"gitlab.com/layunne/users-crud-go/models"
	"gitlab.com/layunne/users-crud-go/repository"
	"net/http"
)

func NewUsersService(
	usersRepository repository.UsersRepository) UsersService {
	return &usersService{usersRepository: usersRepository}
}

type UsersService interface {
	Get(id string) *models.User
	GetAll() []*models.User
	Save(user *models.CreateUser) (*models.User, *errors.Error)
	Update(user *models.User) (*models.User, *errors.Error)
	Delete(id string)
}

type usersService struct {
	usersRepository repository.UsersRepository
}

func (s *usersService) Get(id string) *models.User {
	return s.usersRepository.Get(id)
}

func (s *usersService) GetAll() []*models.User {
	return s.usersRepository.GetAll()
}

func (s *usersService) Save(createUser *models.CreateUser) (*models.User, *errors.Error) {

	if len(createUser.Name) < 4 {
		return nil, errors.New(http.StatusBadRequest, "name needs to be greater 3")
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("uuid error: %v", err.Error()))
	}

	user := &models.User{
		Id:    id.String(),
		Name:  createUser.Name,
		Email: createUser.Email,
	}

	s.usersRepository.Save(user)

	return user, nil
}

func (s *usersService) Update(updateUser *models.User) (*models.User, *errors.Error) {

	user := s.usersRepository.Get(updateUser.Id)

	if user == nil {
		return nil, errors.New(http.StatusNotFound, "user not found for id: "+updateUser.Id)
	}

	if len(user.Name) < 4 {
		return nil, errors.New(http.StatusBadRequest, "name needs to be greater 3")
	}

	user.Name = updateUser.Name
	user.Email = updateUser.Email

	s.usersRepository.Save(user)

	return user, nil
}

func (s *usersService) Delete(id string) {
	s.usersRepository.Delete(id)
}
