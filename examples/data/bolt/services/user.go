
package services

import (
	"fmt"
	"github.com/hidevopsio/hiboot/examples/data/bolt/model"
	"github.com/hidevopsio/hiboot/pkg/starter/data/bolt"
)

type UserService struct {
	// will inject BoltRepository that configured in github.com/hidevopsio/hiboot/pkg/starter/data/bolt
	BoltRepository bolt.Repository `inject:""`
}

func (us *UserService) AddUser(user *model.User) error {
	if us.BoltRepository == nil {
		return fmt.Errorf("repository is not injected")
	}

	return us.BoltRepository.Put(user)
}

func (us *UserService) GetUser(id string) (*model.User, error) {
	if us.BoltRepository == nil {
		return nil, fmt.Errorf("repository is not injected")
	}
	var user model.User
	err := us.BoltRepository.Get(id, &user)
	return &user, err
}

func (us *UserService) DeleteUser(id string) error {
	if us.BoltRepository == nil {
		return fmt.Errorf("repository is not injected")
	}
	return us.BoltRepository.Delete(id, &model.User{})
}

