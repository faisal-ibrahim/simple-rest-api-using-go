package comment

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.Time
}
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComment() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
