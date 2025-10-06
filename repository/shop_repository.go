package repository

import (
	"time"

	"github.com/go-to/bcrd_backend/model"
)

type IShopRepository interface {
	GetShopsTotal(year int32) (int64, error)
	GetShops(time *time.Time, userId string, year int32, keywordParams []string, searchParams []int32, orderParam int32, latitude, longitude float64) (*model.ShopsResult, error)
	GetShop(time *time.Time, userId string, shopId int64) (*model.ShopDetail, error)
}

type ShopRepository struct {
	model model.IShopModel
}

func NewShopRepository(m model.ShopModel) *ShopRepository {
	return &ShopRepository{model: &m}
}

func (r *ShopRepository) GetShopsTotal(year int32) (int64, error) {
	return r.model.CountShopsTotal(year)
}

func (r *ShopRepository) GetShops(time *time.Time, userId string, year int32, keywordParams []string, searchParams []int32, orderParam int32, latitude, longitude float64) (*model.ShopsResult, error) {
	return r.model.FindShops(time, userId, year, keywordParams, searchParams, orderParam, latitude, longitude)
}

func (r *ShopRepository) GetShop(time *time.Time, userId string, shopId int64) (*model.ShopDetail, error) {
	return r.model.FindShop(time, userId, shopId)
}
