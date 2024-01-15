package models

// CommentModel 评论表
type CommentModel struct {
	MODEL

	SubComment         []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"` // 子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID    *uint           `json:"parent_comment_id"`
	Content            string          `gorm:"size:256" json:"content"` // 评论内容
	DiggCount          int             `gorm:"foreignKey:ArticleID" json:"digg_count"`
	CommentCount       int             `gorm:"size:8;default:0;" json:"comment_count"`
	Article            int             `gorm:"foreignKey:ArticleID" json:"-"`
	ArticleID          int             `json:"article_id"`
	User               UserModel       `json:"user"`
	UserID             uint            `json:"user_id"`
}
