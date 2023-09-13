package service

import (
	"go-bbs/app/entity"
	"go-bbs/app/respository"
)

type ThreadService struct {
	threadRepo   *respository.ThreadRepository
	threadEntity *entity.ThreadEntity
	userRepo     *respository.UserRepository
	userEntity   *entity.UserEntity
	forumRepo    *respository.ForumRepository
}

func (serv *ThreadService) List() {

}
