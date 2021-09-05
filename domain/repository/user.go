package repository

import (
	"github.com/ShunyaNagashige/golang-jwt-sample/domain/model"
)

type UserRepository interface {
	// Tokenを作成したいため，*model.Userを返す
	Create(userName, email, password string) (*model.User, error)
}
