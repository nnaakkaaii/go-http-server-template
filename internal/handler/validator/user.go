package validator

import (
	"errors"
	"github.com/nnaakkaaii/go-http-server-template/gen/api"
	"log"
)

type User struct {
	*api.User
}

func (r *User) Register() error {
	log.Println(r.Id, r.Password, r.FirstName, r.LastName, r.Email)
	if r.User == nil {
		return errors.New("request cannot be empty")
	}
	if r.Password == nil || *r.Password == "" {
		return errors.New("password field cannot be empty")
	}
	if r.FirstName == nil || *r.FirstName == "" {
		return errors.New("first name field cannot be empty")
	}
	if r.LastName == nil || *r.LastName == "" {
		return errors.New("last name field cannot be empty")
	}
	if r.Email == nil || *r.Email == "" {
		return errors.New("email field cannot be empty")
	}
	return nil
}

func (r *User) Login() error {
	if r == nil {
		return errors.New("request cannot be empty")
	}
	if r.Email == nil || *r.Email == "" {
		return errors.New("email field cannot be empty")
	}
	if r.Password == nil || *r.Password == "" {
		return errors.New("password field cannot be empty")
	}
	return nil
}
