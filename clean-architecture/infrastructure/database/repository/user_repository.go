package database

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yona3/go-samples/clean-architecture/domain/model"
	"github.com/yona3/go-samples/clean-architecture/ent"
)

type userRepository struct {
	db *ent.Client
}

type UserRepository interface {
	Store(ctx *gin.Context, u *model.User) error
	FindAll(ctx *gin.Context) ([]*model.User, error)
}

func NewUserRepository(db *ent.Client) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Store(ctx *gin.Context, u *model.User) error {
	user, err := r.db.User.Create().
		SetName(u.Name).
		SetAge(u.Age).
		SetEmail(u.Email).
		Save(ctx)
	if err != nil {
		return err
	}

	log.Println("user was created: ", user)

	return nil
}

func (r *userRepository) FindAll(ctx *gin.Context) ([]*model.User, error) {
	d, err := r.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, len(d))
	for i, u := range d {
		users[i] = &model.User{
			ID:    u.ID,
			Name:  u.Name,
			Age:   u.Age,
			Email: u.Email,
		}
	}

	log.Println("get all users: ", users)

	return users, nil
}
