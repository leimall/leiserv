package website

import (
	"leiserv/global"
	website "leiserv/models/website/types"
	"leiserv/utils/upload"
	"mime/multipart"
	"strings"
)

type ProductImgService struct{}

func (e *ProductImgService) Upload(file website.ProductImg) error {
	return global.MALL_DB.Create(&file).Error
}

func (p *ProductImgService) SaveProductImg(header *multipart.FileHeader, noSave string, pdID string) (file website.ProductImg, err error) {
	oss := upload.NewOss()
	filepath, key, uploadErr := oss.UploadFile(header)

	if uploadErr != nil {
		return file, uploadErr
	}
	s := strings.Split(header.Filename, ".")
	f := website.ProductImg{
		ProductID: pdID,
		ImgURL:    filepath,
		Name:      header.Filename,
		Tag:       s[len(s)-1],
		UUID:      key,
	}
	if noSave == "0" {
		return f, p.Upload(f)
	}
	return f, nil
}

func (p *ProductImgService) GetProductImgList(pid string) (list interface{}, total int64, err error) {

	var listObj []website.ProductImg
	err = global.MALL_DB.Where("product_id =?", pid).Order("sort_id asc").Find(&listObj).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return listObj, total, nil
}

func (p *ProductImgService) SetMainImageDB(id int, pid string) (productImg website.ProductImg, err error) {
	var img website.ProductImg
	err = global.MALL_DB.Model(&website.ProductImg{}).Where("product_id =? and id =?", pid, id).Updates(map[string]interface{}{
		"is_main": 1,
		"sort_id": 0,
	}).First(&img).Error
	return img, err
}

func (p *ProductImgService) CancelMainImageDB(pid string) (err error) {
	err = global.MALL_DB.Model(&website.ProductImg{}).Where("product_id =? and is_main = 1", pid).Update("is_main", 0).Error
	return err
}

func (p *ProductImgService) SetSortIDforImageDB(id int, pid string, sort int) (err error) {
	err = global.MALL_DB.Model(&website.ProductImg{}).Where("product_id =? and id =?", pid, id).Update("sort_id", sort).Error
	return err
}

func (p *ProductImgService) GetImageListForProduct(pIDs []string) (imageMap map[string][]website.ProductImg, err error) {
	var image []website.ProductImg
	err = global.MALL_DB.Where("product_id in (?)", pIDs).Order("sort_id asc").Find(&image).Error
	if err != nil {
		return nil, err
	}

	imageMap = make(map[string][]website.ProductImg)
	for _, img := range image {
		imageMap[img.ProductID] = append(imageMap[img.ProductID], img)
	}
	return imageMap, err
}
