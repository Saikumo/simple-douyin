package repository

import (
	"saikumo.org/simple-douyin/src/entity"
	"sync"
)

var (
	userRepository     *UserRepository
	UserRepositoryOnce sync.Once
)

type UserRepository struct {
}

// 单例获取用户DAO
func GetUserRepository() *UserRepository {
	UserRepositoryOnce.Do(func() {
		userRepository = new(UserRepository)
	})
	return userRepository
}

func (repository *UserRepository) IsUserExistByUsername(username string) bool {
	var user entity.User
	DB.Where("username=?", username).Take(&user)
	if user.Id == 0 {
		return false
	}
	return true
}

func (repository *UserRepository) CreateUser(user *entity.User) error {
	return DB.Create(user).Error
}
