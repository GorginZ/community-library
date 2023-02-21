package user

import (
	"errors"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Location string `json:"location"` //library site
}

type UserManager interface {
	GetAll() ([]User, error)
	AddUser(user User) error
}

type UserService struct {
	UserRepository UserManager
}

type UserRepository struct {
	Users []User
}

func (us UserService) GetAll() ([]User, error) {
	return us.UserRepository.GetAll()
}

type UserServiceOption func(*UserService)

func NewUserService(opts ...UserServiceOption) *UserService {
	// defaults
	defaultRepository := UserRepository{}
	us := &UserService{
		UserRepository: &defaultRepository,
	}
	for _, opt := range opts {
		opt(us)
	}
	return us
}

func WithFakeUserRepository() UserServiceOption {
	return func(us *UserService) {
		us.UserRepository = newFakeUserRepository()
	}
}

func (r *FakeUserRepository) AddUser(user User) error {
	r.Users = append(r.Users, user)
	return nil
}

func (r *FakeUserRepository) GetAll() ([]User, error) {
	fmt.Println("fake get all r.Users: ", r.Users)
	return r.Users, nil
}

func newFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{
		Users: []User{
			{
				ID:       1,
				Username: "user1",
				Password: "password1",
				Location: "library1",
			},
		}}
}

type FakeUserRepository struct {
	Users []User
}

func (r *UserRepository) GetAll() ([]User, error) {
	fmt.Println("r.Users: ", r.Users)
	return r.Users, nil
}

func (r *UserRepository) GetUser(id int) (User, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (r *UserRepository) AddUser(user User) error {
	r.Users = append(r.Users, user)
	return nil
}

func (r *UserRepository) UpdateUser(id int, user User) error {
	for i, u := range r.Users {
		if u.ID == id {
			r.Users[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}
