package service

import (
	"context"
	"log"
	"time"
	"xyz-consumer-service/common/config"
	commonErr "xyz-consumer-service/common/error"
	"xyz-consumer-service/modules/consumer_limit/entity"
	"xyz-consumer-service/modules/consumer_limit/internal/repository"
)

type ConsumerLimitService struct {
	cfg                     config.Config
	consumerLimitRepository repository.ConsumerLimitRepositoryUseCase
}

func NewConsumerLimitService(cfg config.Config, consumerLimitRepository repository.ConsumerLimitRepositoryUseCase) *ConsumerLimitService {
	return &ConsumerLimitService{
		cfg:                     cfg,
		consumerLimitRepository: consumerLimitRepository,
	}
}

type ConsumerLimitServiceUseCase interface {
	FindByConsumerId(ctx context.Context, consumerId uint64) ([]*entity.ConsumerLimit, error)
	FindByConsumerIdAndTenor(ctx context.Context, consumerId uint64, tenor uint32) (*entity.ConsumerLimit, error)
	Create(ctx context.Context, consumerId uint64, tenor uint32, limitAmount uint64) (*entity.ConsumerLimit, error)
	UpdateLimitAvailable(ctx context.Context, consumerLimit *entity.ConsumerLimit, transaction uint64) (*entity.ConsumerLimit, error)
}

func (svc *ConsumerLimitService) FindByConsumerId(ctx context.Context, consumerId uint64) ([]*entity.ConsumerLimit, error) {
	res, err := svc.consumerLimitRepository.FindByConsumerId(ctx, consumerId)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitService - FindByConsumerId] Error while find consumer limit by consumer id:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ConsumerLimitService) FindByConsumerIdAndTenor(ctx context.Context, consumerId uint64, tenor uint32) (*entity.ConsumerLimit, error) {
	res, err := svc.consumerLimitRepository.FindByConsumerIdAndTenor(ctx, consumerId, tenor)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitService - FindByConsumerIdAndTenor] Error while find consumer limit by consumer id and tenor:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ConsumerLimitService) Create(ctx context.Context, consumerId uint64, tenor uint32, limitAmount uint64) (*entity.ConsumerLimit, error) {
	consumerLimit := &entity.ConsumerLimit{
		ConsumerId:     consumerId,
		Tenor:          tenor,
		LimitAmount:    limitAmount,
		LimitAvailable: limitAmount,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	res, err := svc.consumerLimitRepository.Create(ctx, consumerLimit)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitService - Create] Error while create consumer limit:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ConsumerLimitService) UpdateLimitAvailable(ctx context.Context, consumerLimit *entity.ConsumerLimit, transaction uint64) (*entity.ConsumerLimit, error) {
	updatedMap := make(map[string]interface{})

	updatedMap["limit_available"] = consumerLimit.LimitAvailable - transaction
	updatedMap["updated_at"] = time.Now()

	res, err := svc.consumerLimitRepository.Update(ctx, consumerLimit, updatedMap)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitService - UpdateLimitAvailable] Error while update consumer limit:", parseError.Message)
		return nil, err
	}

	return res, nil
}
