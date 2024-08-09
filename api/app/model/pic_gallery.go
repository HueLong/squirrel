package model

import (
	"hbbapi/util/db"
)

type PicGalleryModel struct {
}

type PicGallery struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	PicUrl    string `json:"pic_url"`
	CreatedAt int    `json:"created_at"`
}

func (PicGallery) TableName() string {
	return "pic_gallery"
}

func (PicGalleryModel) GetPicList(page, size int) (picList []PicGallery, err error) {
	res := db.Db.Model(&PicGallery{}).
		Where("status = ?", 1).
		Limit(size).
		Offset((page - 1) * size).
		Order("id DESC").
		Find(&picList)
	return picList, res.Error
}
