package servicecontext

import (
	"chatbox/config"
	"chatbox/database"
)

type ServiceContext struct {
	Config    *config.Config
	UserModel *database.UserModel
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: database.NewUserModel(),
	}
}
