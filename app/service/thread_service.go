package service

import (
	"go-bbs/app/http/model"
	"go-bbs/app/respository"
)

type ThreadService struct {
	ForumRepo respository.ThreadRepository
}

func (serv *ThreadService) List(fid int) (threadList []*model.Thread, e error) {
	threadList, e = serv.ForumRepo.ThreadList(fid)
	if e != nil {
		return nil, e
	}
	return threadList, nil
}

func (serv *ThreadService) name() {

}
