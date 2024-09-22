package service

import (
	"context"
	"log"
	"time"
	"xyz-consumer-service/common/config"
	commonErr "xyz-consumer-service/common/error"
	"xyz-consumer-service/common/utils"
	"xyz-consumer-service/modules/consumer/entity"
	"xyz-consumer-service/modules/consumer/internal/repository"
)

type ConsumerService struct {
	cfg                config.Config
	consumerRepository repository.ConsumerRepositoryUseCase
}

func NewConsumerService(cfg config.Config, consumerRepository repository.ConsumerRepositoryUseCase) *ConsumerService {
	return &ConsumerService{
		cfg:                cfg,
		consumerRepository: consumerRepository,
	}
}

type ConsumerServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Consumer, error)
	FindById(ctx context.Context, id uint64) (*entity.Consumer, error)
	Create(ctx context.Context, nik, fullname, legalName, birthPlace, birthDate string, salary uint64, ktp, selfiePhoto string) (*entity.Consumer, error)
	Update(ctx context.Context, id uint64, data *entity.Consumer) (*entity.Consumer, error)
	Delete(ctx context.Context, id uint64) error
}

func (svc *ConsumerService) FindAll(ctx context.Context, req any) ([]*entity.Consumer, error) {
	res, err := svc.consumerRepository.FindAll(ctx, req)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerService - FindAll] Error while find all consumer:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ConsumerService) FindById(ctx context.Context, id uint64) (*entity.Consumer, error) {
	res, err := svc.consumerRepository.FindById(ctx, id)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerService - FindById] Error while find consumer by id:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ConsumerService) Create(ctx context.Context, nik, fullname, legalName, birthPlace, birthDate string, salary uint64, ktp, selfiePhoto string) (*entity.Consumer, error) {
	consumer := &entity.Consumer{
		Nik:         nik,
		Fullname:    fullname,
		LegalName:   legalName,
		BirthPlace:  birthPlace,
		BirthDate:   birthDate,
		Salary:      salary,
		Ktp:         ktp,
		SelfiePhoto: selfiePhoto,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := svc.consumerRepository.Create(ctx, consumer)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerService - Create] Error while create consumer:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ConsumerService) Update(ctx context.Context, id uint64, data *entity.Consumer) (*entity.Consumer, error) {
	consumer, err := svc.consumerRepository.FindById(ctx, id)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerService - Update] Error while find consumer by id:", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})

	utils.AddItemToMap(updatedMap, "nik", data.Nik)
	utils.AddItemToMap(updatedMap, "fullname", data.Fullname)
	utils.AddItemToMap(updatedMap, "legal_name", data.LegalName)
	utils.AddItemToMap(updatedMap, "birth_place", data.BirthPlace)
	utils.AddItemToMap(updatedMap, "birth_date", data.BirthDate)
	utils.AddItemToMap(updatedMap, "salary", data.Salary)
	utils.AddItemToMap(updatedMap, "ktp", data.Ktp)
	utils.AddItemToMap(updatedMap, "selfie_photo", data.SelfiePhoto)

	res, err := svc.consumerRepository.Update(ctx, consumer, updatedMap)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerService - Update] Error while update consumer:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ConsumerService) Delete(ctx context.Context, id uint64) error {
	err := svc.consumerRepository.Delete(ctx, id)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerService - Delete] Error while delete consumer:", parseError.Message)
		return err
	}

	return nil
}
