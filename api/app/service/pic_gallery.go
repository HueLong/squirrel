package service

import (
	"fmt"
	"gorm.io/gorm"
	"hbbapi/app/helper"
	"hbbapi/app/model"
)

type PicGalleryService struct {
	Db *gorm.DB
}

func (s PicGalleryService) GetPicList(page, size int) (list []model.PicGallery) {
	var err error
	list, err = model.PicGalleryModel{}.GetPicList(page, size)
	if err != nil {
		fmt.Println(fmt.Sprintf("GetPicList报错：%s", err.Error()))
		return []model.PicGallery{}
	}
	for k, v := range list {
		v.PicUrl = helper.GetImageLink(v.PicUrl)
		list[k] = v
	}
	return list
}
