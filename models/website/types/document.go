package website

import "leiserv/global"

type Document struct {
	global.DATE_MODEL
	Title   string `gorm:"type:varchar(255);not null;default:'';comment:'文档标题'"`
	Content string `gorm:"type:text;not null;comment:'文档内容'"`
	Status  uint8  `gorm:"type:tinyint unsigned;not null;default:1;comment:'状态 0:deleted, 1:active'"`
}

func (Document) TableName() string {
	return "document"
}
