package models

// MessageModel 记录消息
type MessageModel struct {
	MODEL

	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"`               // 发送人 ID，使用gorm库的primaryKey标签指定为主键，json标签指定在JSON序列化时字段名称为send_user_id
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"send_user_model"` // 发送人的用户模型，使用gorm库的foreignKey标签指定外键关联到UserModel结构体，json标签指定在JSON序列化时字段名称为send_user_model
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"`           // 发送人的昵称，使用gorm库的size标签指定字段长度为42，json标签指定在JSON序列化时字段名称为send_user_nick_name
	SendUserAvatar   string    `json:"send_user_avatar"`                             // 发送人的头像，json标签指定在JSON序列化时字段名称为send_user_avatar

	RevUserID       uint      `gorm:"primaryKey" json:"rev_user_id"`              // 接收人 ID，使用gorm库的primaryKey标签指定为主键，json标签指定在JSON序列化时字段名称为rev_user_id
	RevUserModel    UserModel `gorm:"foreignKey:RevUserID" json:"rev_user_model"` // 接收人的用户模型，使用gorm库的foreignKey标签指定外键关联到UserModel结构体，json标签指定在JSON序列化时字段名称为rev_user_model
	RevUserNickName string    `gorm:"size:42" json:"rev_user_nick_name"`          // 接收人的昵称，使用gorm库的size标签指定字段长度为42，json标签指定在JSON序列化时字段名称为rev_user_nick_name
	RevUserAvatar   string    `json:"rev_user_avatar"`                            // 接收人的头像，json标签指定在JSON序列化时字段名称为rev_user_avatar
	IsRead          bool      `gorm:"default:false" json:"is_read"`               // 消息是否已读，默认为false，使用gorm库的default标签指定默认值为false，json标签指定在JSON序列化时字段名称为is_read
	Content         string    `json:"content"`                                    // 消息内容，json标签指定在JSON序列化时字段名称为content
}
