package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Stamp struct {
	gorm.Model
	UserID        string
	ShopID        int64
	NumberOfTimes int32
}

type StampDetail struct {
	gorm.Model
	UserID    string
	ShopID    int64
	StampedAt time.Time
}

func (Stamp) TableName() string {
	return "stamps"
}

func (StampDetail) TableName() string {
	return "stamps_detail"
}

type IStampModel interface {
	AddStamp(time *time.Time, userId string, shopId int64) (int32, error)
	DeleteStamp(userId string, shopId int64) (int32, error)
}

type StampModel struct {
	db DB
}

func NewStampModel(db DB) *StampModel {
	return &StampModel{db: db}
}

func (m *StampModel) AddStamp(time *time.Time, userId string, shopId int64) (int32, error) {
	recordNum := int64(0)
	err := m.db.Conn.Transaction(func(tx *gorm.DB) error {
		// stamps_detailテーブルにレコードを登録
		if err := tx.Create(&StampDetail{
			UserID:    userId,
			ShopID:    shopId,
			StampedAt: *time,
		}).Error; err != nil {
			return err
		}

		// stamps_detailテーブルのレコード数の情報を取得
		if err := tx.Model(&StampDetail{}).
			Where("user_id = ? and shop_id = ?", userId, shopId).
			Count(&recordNum).Error; err != nil {
			return err
		}

		// stampsテーブルの既存レコードを取得
		stamp := &Stamp{}
		if err := tx.Where("user_id = ? and shop_id = ?", userId, shopId).
			First(stamp).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// stampsテーブルにレコードを登録・更新
		stamp.UserID = userId
		stamp.ShopID = shopId
		stamp.NumberOfTimes = int32(recordNum)
		if err := tx.Save(stamp).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return int32(recordNum), nil
}

func (m *StampModel) DeleteStamp(userId string, shopId int64) (int32, error) {
	recordNum := int64(0)
	err := m.db.Conn.Transaction(func(tx *gorm.DB) error {
		// stamps_detailテーブルから最新のレコードを1件取得
		stampDetail := &StampDetail{}
		err := tx.Where("user_id = ? and shop_id = ?", userId, shopId).
			Order("stamped_at desc").
			First(stampDetail).Error
		// 削除対象のレコードが存在しない場合は正常終了
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else if err != nil {
			return err
		}

		// stamps_detailテーブルからレコードを削除
		if err := tx.Where("id = ?", stampDetail.ID).
			Delete(&StampDetail{}).Error; err != nil {
			return err
		}

		// stamps_detailテーブルのレコード数の情報を取得
		if err := tx.Model(&StampDetail{}).
			Where("user_id = ? and shop_id = ?", userId, shopId).
			Count(&recordNum).Error; err != nil {
			return err
		}

		// stampsテーブルのレコードを更新
		stamp := &Stamp{}
		if err := tx.Where("user_id = ? and shop_id = ?", userId, shopId).
			First(stamp).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// stampsテーブルのレコードを更新
		stamp.UserID = userId
		stamp.ShopID = shopId
		stamp.NumberOfTimes = int32(recordNum)
		if err := tx.Save(stamp).Error; err != nil {
			return err
		}

		// stamps_detailテーブルのレコードが0件の場合はstampsテーブルのレコードを削除
		if recordNum == 0 {
			tx.Where("user_id = ? and shop_id = ?", userId, shopId).
				Delete(&Stamp{})
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return int32(recordNum), nil
}
