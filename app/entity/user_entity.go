package entity

import "go-bbs/app/http/model"

type UserEntity struct {
	*model.User
	Group *model.Group
}
