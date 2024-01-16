package models

// CommentModel 评论表
type CommentModel struct {
	MODEL

	SubComment         []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  // 子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` // 父评论模型
	ParentCommentID    *uint           `json:"parent_comment_id"`                               // 父评论ID
	Content            string          `gorm:"size:256" json:"content"`                         // 评论内容
	DiggCount          int             `gorm:"foreignKey:ArticleID" json:"digg_count"`          // 点赞数量
	CommentCount       int             `gorm:"size:8;default:0;" json:"comment_count"`          // 评论数量
	Article            int             `gorm:"foreignKey:ArticleID" json:"-"`                   // 文章
	ArticleID          int             `json:"article_id"`                                      // 文章ID
	User               UserModel       `json:"user"`                                            // 用户
	UserID             uint            `json:"user_id"`                                         // 用户ID
}
