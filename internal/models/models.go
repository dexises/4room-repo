package models

// import (
// 	"database/sql"
// 	"errors"
// )

// var (
// 	ErrRecordNotFound = errors.New("record not found")
// 	ErrEditConflict   = errors.New("edit conflict")
// )

// type Models struct {
// 	Post interface {
// 		Insert(post *Post) error
// 		Get(id int64) (*Post, error)
// 		Update(post *Post) error
// 		Delete(id int64) error
// 		GetAll(subject string, category []string, filters Filters) ([]*Post, Metadata, error)
// 	}
// 	Users       UserModel
// 	Tokens      TokenModel
// 	Permissions PermissionModel
// }

// func NewMockModels() Models {
// 	return Models{
// 		Posts: MockPostModel{},
// 	}
// }

// func NewModels(db *sql.DB) Models {
// 	return Models{
// 		Posts: PostModel{DB: db},
// 	}
// }
