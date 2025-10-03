package usecase

import (
	"github.com/go-to/bcrd_backend/repository"
	"github.com/go-to/bcrd_backend/usecase/input"
	"github.com/go-to/bcrd_backend/usecase/output"
	"github.com/go-to/bcrd_protobuf/pb"
)

type IStampUsecase interface {
	AddStamp(in *input.StampInput) (*output.StampOutput, error)
	DeleteStamp(in *input.StampInput) (*output.StampOutput, error)
}

type StampUsecase struct {
	config repository.IConfigRepository
	stamp  repository.IStampRepository
}

func NewStampUseCase(config repository.ConfigRepository, stamp repository.StampRepository) *StampUsecase {
	return &StampUsecase{
		config: &config,
		stamp:  &stamp,
	}
}

func (u *StampUsecase) AddStamp(in *input.StampInput) (*output.StampOutput, error) {
	userId := in.StampRequest.GetUserId()
	shopId := in.StampRequest.GetShopId()

	now, err := u.config.GetTime()
	if err != nil {
		return &output.StampOutput{}, err
	}

	stampNum, err := u.stamp.AddStamp(&now, userId, shopId)
	if err != nil {
		return &output.StampOutput{}, err
	}

	return &output.StampOutput{
		StampResponse: pb.StampResponse{
			NumberOfTimes: stampNum,
		},
	}, nil
}

func (u *StampUsecase) DeleteStamp(in *input.StampInput) (*output.StampOutput, error) {
	userId := in.StampRequest.GetUserId()
	shopId := in.StampRequest.GetShopId()

	stampNum, err := u.stamp.DeleteStamp(userId, shopId)
	if err != nil {
		return &output.StampOutput{}, err
	}

	return &output.StampOutput{
		StampResponse: pb.StampResponse{
			NumberOfTimes: stampNum,
		},
	}, nil
}
