package api

type RouterGroup struct {
	AuthRouter
	UserRouter
	ForumRouter
	ThreadRouter
	PostRouter
	KaDaoDataRouter
}
