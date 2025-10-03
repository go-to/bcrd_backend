package usecase

import (
	"fmt"
	"slices"
	"strings"

	"github.com/go-to/bcrd_backend/repository"
	"github.com/go-to/bcrd_backend/usecase/input"
	"github.com/go-to/bcrd_backend/usecase/output"
	"github.com/go-to/bcrd_backend/util"
	"github.com/go-to/bcrd_protobuf/pb"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// TODO 見直し
const defaultYear = int32(2025)
const initialLatitude = 35.64691938518296
const initialLongitude = 139.71008179999998

type IShopUsecase interface {
	getDefaultYear() (int32, error)
	GetShopsTotal(in *input.ShopsTotalInput) (*output.ShopsTotalOutput, error)
	GetShops(in *input.ShopsInput) (*output.ShopsOutput, error)
	GetShop(in *input.ShopInput) (*output.ShopOutput, error)
}

type ShopUsecase struct {
	config repository.IConfigRepository
	event  repository.IEventRepository
	shop   repository.IShopRepository
}

func NewShopUseCase(config repository.ConfigRepository, event repository.EventRepository, shop repository.ShopRepository) *ShopUsecase {
	return &ShopUsecase{
		config: &config,
		event:  &event,
		shop:   &shop,
	}
}

func (u *ShopUsecase) getDefaultYear() (int32, error) {
	// TODO 見直し
	//now, err := u.config.GetTime()
	//if err != nil {
	//	return 0, err
	//}
	//return int32(now.Year()), nil
	return defaultYear, nil
}

func (u *ShopUsecase) GetShopsTotal(in *input.ShopsTotalInput) (*output.ShopsTotalOutput, error) {
	year := in.ShopsTotalRequest.GetYear()
	if year == 0 {
		var err error
		year, err = u.getDefaultYear()
		if err != nil {
			return &output.ShopsTotalOutput{}, err
		}
	}

	shopsTotal, err := u.shop.GetShopsTotal(year)
	if err != nil {
		return &output.ShopsTotalOutput{}, err
	}

	return &output.ShopsTotalOutput{
		ShopsTotalResponse: pb.ShopsTotalResponse{
			TotalNum: shopsTotal,
		},
	}, nil
}

func (u *ShopUsecase) GetShops(in *input.ShopsInput) (*output.ShopsOutput, error) {
	year := in.ShopsRequest.GetYear()
	if year == 0 {
		var err error
		year, err = u.getDefaultYear()
		if err != nil {
			return &output.ShopsOutput{}, err
		}
	}
	userId := in.ShopsRequest.GetUserId()
	searchTypes := in.ShopsRequest.GetSearchTypes()
	keywords := in.ShopsRequest.GetKeyword()
	sortOrder := in.ShopsRequest.GetSortOrder()
	latitude := in.ShopsRequest.GetLatitude()
	longitude := in.ShopsRequest.GetLongitude()

	if latitude == 0 && longitude == 0 {
		latitude = initialLatitude
		longitude = initialLongitude
	}

	var searchParams []int32
	for _, value := range searchTypes {
		v := int32(value)
		if slices.Contains(searchParams, v) {
			continue
		}
		if key, exists := pb.SearchType_name[v]; exists {
			searchParams = append(searchParams, pb.SearchType_value[key])
		}
	}
	// 検索キーワードの整形
	keywordParams := strings.Fields(keywords)

	// ソート順
	orderParam := int32(pb.SortOrderType_SORT_ORDER_NO)
	if key, exists := pb.SortOrderType_name[int32(sortOrder)]; exists {
		orderParam = pb.SortOrderType_value[key]
	}

	now, err := u.config.GetTime()
	if err != nil {
		return &output.ShopsOutput{}, err
	}

	shops, err := u.shop.GetShops(&now, userId, year, keywordParams, searchParams, orderParam, latitude, longitude)
	if err != nil {
		return &output.ShopsOutput{}, err
	}

	var outputShops []*pb.Shop
	var latLonList []string

	for _, v := range *shops {
		inCurrentSales := true
		if len(v.StartTime) == 0 || len(v.EndTime) == 0 {
			inCurrentSales = false
		}
		// 緯度経度が同じ場合は、重なり防止のためにマーカーの位置をずらす
		lat := v.Latitude
		lon := v.Longitude
		latLon := fmt.Sprintf("%f,%f", lat, lon)
		if slices.Contains(latLonList, latLon) {
			lat += 0.00002
			lon += 0.00002
		}
		latLonList = append(latLonList, latLon)

		// 距離（1,000m以上の場合は単位をkmに変更する）
		distance := util.FormatDistance(v.Distance)

		isStamped := false
		if v.NumberOfTimes > 0 {
			isStamped = true
		}

		outputShops = append(outputShops, &pb.Shop{
			Id:                         v.ID,
			EventId:                    v.EventID,
			Year:                       v.Year,
			CategoryId:                 pb.CategoryType(v.CategoryID),
			CategoryName:               v.CategoryName,
			No:                         v.No,
			ShopName:                   v.ShopName,
			MenuName:                   v.MenuName,
			MenuImageUrl:               v.MenuImageUrl,
			Phone:                      v.Phone,
			Address:                    v.Address,
			BusinessDays:               v.BusinessDays,
			RegularHoliday:             v.RegularHoliday,
			BusinessHours:              v.BusinessHours,
			ChargePrice:                v.ChargePrice,
			NormalizedChargePrice:      v.NormalizedChargePrice,
			SinglePrice:                v.SinglePrice,
			NormalizedSinglePrice:      v.NormalizedSinglePrice,
			SetPrice:                   v.SetPrice,
			NormalizedSetPrice:         v.NormalizedSetPrice,
			BeerType:                   v.BeerType,
			NeedsReservation:           v.NeedsReservation,
			NormalizedNeedsReservation: v.NormalizedNeedsReservation,
			UseHachipay:                v.UseHachipay,
			NormalizedUseHachipay:      v.NormalizedUseHachipay,
			IsOpenHoliday:              v.IsOpenHoliday,
			IsIrregularHoliday:         v.IsIrregularHoliday,
			Latitude:                   lat,
			Longitude:                  lon,
			Distance:                   distance,
			WeekNumber:                 v.WeekNumber,
			DayOfWeek:                  int32(v.DayOfWeek),
			StartTime:                  v.StartTime,
			EndTime:                    v.EndTime,
			IsHoliday:                  v.IsHoliday,
			InCurrentSales:             inCurrentSales,
			NumberOfTimes:              v.NumberOfTimes,
			IsStamped:                  isStamped,
		})
	}

	fmt.Println("")
	fmt.Printf("lat: %+v, long: %+v\n", latitude, longitude)
	if len(outputShops) > 0 {
		fmt.Printf("dist: %+v\n", outputShops[0].Distance)
	}

	return &output.ShopsOutput{
		ShopsResponse: pb.ShopsResponse{
			Shops: outputShops,
		},
	}, nil
}

func (u *ShopUsecase) GetShop(in *input.ShopInput) (*output.ShopOutput, error) {
	userId := in.ShopRequest.GetUserId()
	shopId := in.ShopRequest.GetShopId()

	now, err := u.config.GetTime()
	if err != nil {
		return &output.ShopOutput{}, err
	}

	shop, err := u.shop.GetShop(&now, userId, shopId)
	if err != nil {
		return &output.ShopOutput{}, err
	}

	outputShop := &pb.Shop{}
	if &shop != nil {
		inCurrentSales := true
		if len(shop.StartTime) == 0 || len(shop.EndTime) == 0 {
			inCurrentSales = false
		}

		// 距離
		fmtX := message.NewPrinter(language.Japanese)
		distance := fmtX.Sprintf("%dm", int(shop.Distance))

		isStamped := false
		if shop.NumberOfTimes > 0 {
			isStamped = true
		}

		outputShop = &pb.Shop{
			Id:                         shop.ID,
			EventId:                    shop.EventID,
			Year:                       shop.Year,
			CategoryId:                 pb.CategoryType(shop.CategoryID),
			CategoryName:               shop.CategoryName,
			No:                         shop.No,
			ShopName:                   shop.ShopName,
			MenuName:                   shop.MenuName,
			MenuImageUrl:               shop.MenuImageUrl,
			Phone:                      shop.Phone,
			Address:                    shop.Address,
			BusinessDays:               shop.BusinessDays,
			RegularHoliday:             shop.RegularHoliday,
			BusinessHours:              shop.BusinessHours,
			ChargePrice:                shop.ChargePrice,
			NormalizedChargePrice:      shop.NormalizedChargePrice,
			SinglePrice:                shop.SinglePrice,
			NormalizedSinglePrice:      shop.NormalizedSinglePrice,
			SetPrice:                   shop.SetPrice,
			NormalizedSetPrice:         shop.NormalizedSetPrice,
			BeerType:                   shop.BeerType,
			NeedsReservation:           shop.NeedsReservation,
			NormalizedNeedsReservation: shop.NormalizedNeedsReservation,
			UseHachipay:                shop.UseHachipay,
			NormalizedUseHachipay:      shop.NormalizedUseHachipay,
			IsOpenHoliday:              shop.IsOpenHoliday,
			IsIrregularHoliday:         shop.IsIrregularHoliday,
			Latitude:                   shop.Latitude,
			Longitude:                  shop.Longitude,
			Distance:                   distance,
			WeekNumber:                 shop.WeekNumber,
			DayOfWeek:                  int32(shop.DayOfWeek),
			StartTime:                  shop.StartTime,
			EndTime:                    shop.EndTime,
			IsHoliday:                  shop.IsHoliday,
			InCurrentSales:             inCurrentSales,
			NumberOfTimes:              shop.NumberOfTimes,
			IsStamped:                  isStamped,
		}
	}

	isEventPeriod := false
	isCheckEventPeriod, err := u.config.IsCheckEventPeriod()
	if err != nil {
		return &output.ShopOutput{}, err
	}
	if isCheckEventPeriod {
		activeEvent, err := u.event.GetActiveEvents(&now)
		if err != nil {
			return &output.ShopOutput{}, err
		}
		if activeEvent.ID != 0 {
			isEventPeriod = true
		}
	} else {
		isEventPeriod = true
	}

	return &output.ShopOutput{
		ShopResponse: pb.ShopResponse{
			Shop:          outputShop,
			IsEventPeriod: isEventPeriod,
		},
	}, nil
}
