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
	AddUser(u User) error
}

type UserService struct {
	UserRepository UserManager
}

func NewUserService(opts ...UserServiceOption) *UserService {
	defaultRepository := UserRepository{}
	us := &UserService{
		UserRepository: defaultRepository,
	}
	for _, opt := range opts {
		opt(us)
	}
	return us
}

type UserRepository struct {
}

type UserServiceOption func(*UserService)

func WithFakeUserRepository() UserServiceOption {
	return func(us *UserService) {
		us.UserRepository = newFakeUserRepository()
	}
}

func (r *UserRepository) GetAll() ([]User, error) {
	fmt.Println("REAL UserRepository.GetAll")
	return nil, errors.New("not implemented")
}

type FakeUserRepository struct {
	Users []User
}

func newFakeUserRepository() FakeUserRepository {
	ur := FakeUserRepository{}
	return ur
}

func (ur *FakeUserRepository) GetAll() ([]User, error) {
	fmt.Println("FakeUserRepository.GetAll")
	if ur.Users != nil {
		return ur.Users, nil
	}
	return ur.Users, nil
}

func (r *FakeUserRepository) AddUser(u User) error {
	r.Users = append(r.Users, u)
	fmt.Println("users: ", r.Users)
	return nil
}

func (r *UserRepository) AddUser(u User) error {
	return errors.New("not implemented")
}
