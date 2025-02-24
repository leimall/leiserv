package website

import (
	"fmt"
	"leiserv/global"
	website "leiserv/models/website/types"
)

type DocumentService struct{}

func (d *DocumentService) GetDocumentByTitleFromDB(title string) (doc website.Document, err error) {
	err = global.MALL_DB.Where("title LIKE?", "%"+title+"%").First(&doc).Error
	fmt.Println("aaal", title, doc.Title)
	return doc, err
}
