package entity

import (
	"go-bbs/app/http/model"
	"sync"
)

type UserEntity struct {
	model.User
	Group *model.Group //用户组
	mu    sync.Mutex
}

// Identity 唯一
func (u *UserEntity) Identity() int {
	return u.Uid
}

// ChangeCredits 经验值修改
func (u *UserEntity) ChangeCredits(n int) {
	u.mu.Lock()
	u.SetCredits(n)
	u.mu.Unlock()
}

// ChangeGolds 经验值修改
func (u *UserEntity) ChangeGolds(n int) {
	u.mu.Lock()
	u.SetGolds(n)
	u.mu.Unlock()
}

// ChangeRmbs 经验值修改
func (u *UserEntity) ChangeRmbs(n int) {
	u.SetRmbs(n)
}
