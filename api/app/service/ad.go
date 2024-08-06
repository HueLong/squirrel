package service

import (
	"fmt"
	"gorm.io/gorm"
	"hbbapi/app/enum/cache_enum"
	"hbbapi/app/model"
	"hbbapi/util/cache"
)

type AdService struct {
	Db *gorm.DB
}

// getAdByName 通过广告位名称获取广告
// 缓存10分钟
func (s AdService) getAdByName(name string) []model.AdBase {
	var ads []model.AdBase
	key := fmt.Sprintf(cache_enum.AdByName, name)
	err := cache.GetInstance().GetJson(key, &ads)
	if err != nil {
		ads = model.AdModel{}.GetAdByName(name)
		_ = cache.GetInstance().SetJson(key, ads, 600)
	}
	if ads == nil {
		return []model.AdBase{}
	}
	return ads
}

// getAdByShortName 通过广告位标识获取广告
// 缓存10分钟
func (s AdService) getAdByShortName(shortName string) []model.AdBase {
	var ads []model.AdBase
	key := fmt.Sprintf(cache_enum.AdByShortName, shortName)
	err := cache.GetInstance().GetJson(key, &ads)
	if err != nil {
		ads = model.AdModel{}.GetAdByShortName(shortName)
		_ = cache.GetInstance().SetJson(key, ads, 600)
	}
	if ads == nil {
		return []model.AdBase{}
	}
	return ads
}

// GetAdByShowPosition 根据展示位置标识获取广告信息
func (s AdService) GetAdByShowPosition(showPosition string) model.AdBase {
	var ad model.AdBase
	key := fmt.Sprintf(cache_enum.AdByShowPosition, showPosition)
	err := cache.GetInstance().GetJson(key, &ad)
	if err != nil {
		adList := model.AdModel{}.GetAdByShowPosition(showPosition)
		if len(adList) > 0 {
			ad = adList[0]
		}
		_ = cache.GetInstance().SetJson(key, ad, 600)
	}
	return ad
}

func (s AdService) GetAdList(showPosition string) []model.AdBase {
	var adList []model.AdBase
	key := fmt.Sprintf(cache_enum.AdByShowPosition, showPosition)
	err := cache.GetInstance().GetJson(key, &adList)
	if err != nil {
		adList = model.AdModel{}.GetAdByShowPosition(showPosition)
		_ = cache.GetInstance().SetJson(key, adList, 600)
	}
	return adList
}

// getAdByIds 获取广告信息
func (s AdService) getAdByIds(ids ...int) map[int]model.Ad {
	var adMap = make(map[int]model.Ad)
	var notCachedIds []int
	for _, id := range ids {
		cacheKey := fmt.Sprintf(cache_enum.Ad, id)
		var ad model.Ad
		err := cache.GetInstance().GetJson(cacheKey, &ad)
		if err != nil {
			notCachedIds = append(notCachedIds, id)
		} else {
			adMap[id] = ad
		}
	}
	if len(notCachedIds) > 0 {
		ads := model.AdModel{}.GetAdByIds(notCachedIds...)
		for _, value := range ads {
			cacheKey := fmt.Sprintf(cache_enum.Ad, value.Id)
			_ = cache.GetInstance().SetJson(cacheKey, value, -1)
			adMap[value.Id] = value
		}
	}
	return adMap
}

// 获取商品详情页广告
func (s AdService) getGoodsDetailAd() (top []model.Ad, bottom []model.Ad) {
	topAdName := "shop1"
	bottomAdName := "shop2"
	topIds := s.getAdPositionIdByName(topAdName)
	if len(topIds) > 0 {
		topMap := s.getAdByIds(topIds...)
		for _, value := range topMap {
			top = append(top, value)
		}
	}
	bottomIds := s.getAdPositionIdByName(bottomAdName)
	if len(bottomIds) > 0 {
		bottomMap := s.getAdByIds(bottomIds...)
		for _, value := range bottomMap {
			bottom = append(bottom, value)
		}
	}
	return
}

// TODO 修改后台广告编辑
func (s AdService) getAdPositionIdByName(name string) []int {
	var ids []int
	key := fmt.Sprintf(cache_enum.AdByName, name)
	err := cache.GetInstance().GetJson(key, &ids)
	if err != nil {
		ids = model.AdModel{}.GetAdPositionIdByName(name)
		_ = cache.GetInstance().SetJson(key, ids, 86400)
	}
	if ids == nil {
		return []int{}
	}
	return ids
}

// getAdsByTagId 获取标签的广告
func (s AdService) getAdsByTagId(tagId int) []model.AdBase {
	var ads []model.AdBase
	key := fmt.Sprintf(cache_enum.AdsByTagId, tagId)
	err := cache.GetInstance().GetJson(key, &ads)
	if err != nil {
		ads = model.AdModel{}.GetAdsByTagId(tagId)
		_ = cache.GetInstance().SetJson(key, ads, 600)
	}
	if ads == nil {
		return []model.AdBase{}
	}
	return ads
}
