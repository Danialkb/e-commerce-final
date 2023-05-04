package comments

type CommentServiceInterface interface {
	CreateComment(*Comment) error
	GetCommentById(uint) (*Comment, error)
	DeleteComment(uint) error
}

type CommentServiceV1 struct {
	commentRepos CommentRepositoryInterface
}

func NewCommentService() CommentServiceV1 {
	return CommentServiceV1{commentRepos: NewCommentRepository()}
}

func (c CommentServiceV1) CreateComment(comment *Comment) error {
	return c.commentRepos.CreateComment(comment)
}

func (c CommentServiceV1) GetCommentById(id uint) (*Comment, error) {
	return c.commentRepos.GetCommentById(id)
}

func (c CommentServiceV1) DeleteComment(id uint) error {
	return c.commentRepos.DeleteComment(id)
}
