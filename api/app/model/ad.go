package model

import (
	"hbbapi/app/enum"
	"hbbapi/util/db"
	"time"
)

type Ad struct {
	Id           int       `json:"id"`
	AdName       string    `json:"ad_name"`
	AdPositionId int       `json:"ad_position_id"`
	AdLink       string    `json:"ad_link"`
	Tags         string    `json:"tags"`
	AdCode       string    `json:"ad_code"`
	IsShow       string    `json:"is_show"`
	Sort         string    `json:"sort"`
	IsTarget     int8      `json:"is_target"`
	BgColor      string    `json:"bg_color" gorm:"column:bgcolor"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	CreatedAt    time.Time `json:"created_at"`
	AppLink      string    `json:"app_link"`
	AppLinkType  int8      `json:"app_link_type"`
	TopicTags    string    `json:"topic_tags"`
	Type         int8      `json:"type"`
	ProgramId    int       `json:"program_id"`
	ShowPosition string    `json:"show_position"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
}

type AdBase struct {
	Id     int    `json:"id"`
	AdName string `json:"ad_name"`
	AdLink string `json:"ad_link"`
	AdCode string `json:"ad_code"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type AdModel struct {
}

type AdPosition struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Desc      string `json:"desc"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsShow    int    `json:"is_show"`
	Type      int    `json:"type"`
}

func (Ad) TableName() string {
	return "pc_ad"
}

func (AdPosition) TableName() string {
	return "pc_ad_position"
}

func (AdModel) GetAdByName(name string) []AdBase {
	var ads []AdBase
	db.Db.Table("pc_ad a").
		Joins("inner join pc_ad_position p ON a.ad_position_id=p.id").
		Where("p.name=? and a.is_show=1 and p.is_show=1", name).
		Select("a.id,a.ad_link,a.ad_code,a.ad_name,a.width,a.height").
		Find(&ads)
	return ads
}

func (AdModel) GetAdByShortName(shortName string) []AdBase {
	var ads []AdBase
	db.Db.Table("pc_ad a").
		Joins("inner join pc_ad_position p ON a.ad_position_id=p.id").
		Where("p.short_name = ? and a.is_show = 1 and p.is_show = 1", shortName).
		Select("a.id,a.ad_link,a.ad_code,a.ad_name,a.width,a.height").
		Find(&ads)
	return ads
}

// GetAdByShowPosition 根据广告展示位置标识获取广告信息
func (AdModel) GetAdByShowPosition(showPosition string) []AdBase {
	var ad []AdBase
	db.Db.Model(&Ad{}).
		Where("show_position=? and is_show=1", showPosition).
		Order("sort DESC,id DESC").
		Find(&ad)
	return ad
}

// GetAdByPositionIds 根据广告展示位置标识获取广告信息
func (AdModel) GetAdByPositionIds(positionId ...string) AdBase {
	var ad AdBase
	db.Db.Model(&Ad{}).
		Where("ad_position_id IN ? and is_show=1 and type=2", positionId).
		Order("id DESC").
		Find(&ad)
	return ad
}

// GetAdsByTagId 根据标签获取广告列表
func (AdModel) GetAdsByTagId(tagId int) []AdBase {
	var ads []AdBase
	db.Db.Model(&Ad{}).
		Select("id,ad_link,ad_code,width,height,ad_name").
		Where("ad_position_id = 6 AND is_show = ? AND (FIND_IN_SET(?,`tags`) OR tags=?)", enum.AdIsShowTrue, tagId, "").
		Order("sort DESC,id DESC").
		Find(&ads)
	return ads
}

// GetAdPositionIdByName 根据名称获取对应的广告位置
func (AdModel) GetAdPositionIdByName(name string) []int {
	var ids []int
	db.Db.Model(&AdPosition{}).Select("id").
		Where("short_name=? and is_show=1", name).
		Find(&ids)
	return ids
}

// GetAdByIds 通过一组ID获取Ad 信息
func (m AdModel) GetAdByIds(id ...int) []Ad {
	var ads []Ad
	ads = []Ad{}
	db.Db.Where("id IN ? and is_show=1", id).Find(&ads)
	return ads
}
