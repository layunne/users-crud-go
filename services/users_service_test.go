package services

import (
	"gitlab.com/layunne/users-crud-go/errors"
	"gitlab.com/layunne/users-crud-go/models"
	"gitlab.com/layunne/users-crud-go/repository"
	"reflect"
	"sort"
	"testing"
)

func Test_usersService_Get(t *testing.T) {

	user1 := &models.User{
		Id:    "id-1",
		Name:  "Name 1",
		Email: "email1@aa.com",
	}

	usersRepository := repository.NewUsersRepository()
	usersRepository.Save(user1)

	type fields struct {
		usersRepository repository.UsersRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.User
	}{
		{
			name:   "get1",
			fields: fields{usersRepository},
			args:   args{
				id: user1.Id,
			},
			want:   user1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersService{
				usersRepository: tt.fields.usersRepository,
			}
			if got := s.Get(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_GetAll(t *testing.T) {

	user1 := &models.User{
		Id:    "id-1",
		Name:  "Name 1",
		Email: "email1@aa.com",
	}

	user2 := &models.User{
		Id:    "id-2",
		Name:  "Name 2",
		Email: "email2@aa.com",
	}

	user3 := &models.User{
		Id:    "id-3",
		Name:  "Name 3",
		Email: "email3@aa.com",
	}


	usersRepository := repository.NewUsersRepository()
	usersRepository.Save(user1)
	usersRepository.Save(user2)
	usersRepository.Save(user3)

	type fields struct {
		usersRepository repository.UsersRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []*models.User
	}{
		{
			name:   "getAll1",
			fields: fields{usersRepository},
			want: []*models.User{
				user1,user2,user3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersService{
				usersRepository: tt.fields.usersRepository,
			}

			sortUsers(tt.want)

			if got := s.GetAll(); !reflect.DeepEqual(sortUsers(got), sortUsers(tt.want)) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Save(t *testing.T) {

	type fields struct {
		usersRepository repository.UsersRepository
	}
	type args struct {
		createUser *models.CreateUser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.User
		want1  *errors.Error
	}{
		{
			name:   "save1",
			fields: fields{repository.NewUsersRepository()},
			args:   args{
				createUser: &models.CreateUser{
					Name: "Name aa",
					Email: "aaa@aa.com",
				},
			},
			want:   &models.User{
				Name:  "Name aa",
				Email: "aaa@aa.com",
			},
			want1:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, got1 := s.Save(tt.args.createUser)

			if got != nil {
				tt.want.Id = got.Id
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Save() got1 = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}

func Test_usersService_Update(t *testing.T) {

	user1 := &models.User{
		Id:    "id-1",
		Name:  "Name 1",
		Email: "email1@aa.com",
	}

	usersRepository := repository.NewUsersRepository()
	usersRepository.Save(user1)

	user1.Name = "New name"

	type fields struct {
		usersRepository repository.UsersRepository
	}
	type args struct {
		updateUser *models.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.User
		want1  *errors.Error
	}{
		{
			name:   "Update1",
			fields: fields{usersRepository},
			args:   args{
				user1,
			},
			want:   user1,
			want1:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersService{
				usersRepository: tt.fields.usersRepository,
			}
			got, got1 := s.Update(tt.args.updateUser)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Update() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_usersService_Delete(t *testing.T) {

	user1 := &models.User{
		Id:    "id-1",
		Name:  "Name 1",
		Email: "email1@aa.com",
	}

	usersRepository := repository.NewUsersRepository()
	usersRepository.Save(user1)


	type fields struct {
		usersRepository repository.UsersRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "get1",
			fields: fields{usersRepository},
			args:   args{
				id: user1.Id,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &usersService{
				usersRepository: tt.fields.usersRepository,
			}
			s.Delete(tt.args.id)

			var want *models.User = nil

			if got := s.Get(tt.args.id); !reflect.DeepEqual(got, want) {
				t.Errorf("Delete() got = %v, want %v", got, want)
			}
		})
	}
}

func sortUsers(users []*models.User) []*models.User {
	sort.Slice(users[:], func(i, j int) bool {
		a := users[i]
		b := users[j]
		return a.Id < b.Id
	})
	return users
}
