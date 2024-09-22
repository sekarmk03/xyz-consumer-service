package repository

import (
	"context"
	"log"
	"time"
	"xyz-consumer-service/modules/consumer_limit/entity"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

type ConsumerLimitRepository struct {
	db *gorm.DB
}

func NewConsumerLimitRepository(db *gorm.DB) *ConsumerLimitRepository {
	return &ConsumerLimitRepository{
		db: db,
	}
}

type ConsumerLimitRepositoryUseCase interface {
	FindByConsumerId(ctx context.Context, consumerId uint64) ([]*entity.ConsumerLimit, error)
	FindByConsumerIdAndTenor(ctx context.Context, consumerId uint64, tenor uint32) (*entity.ConsumerLimit, error)
	Create(ctx context.Context, req *entity.ConsumerLimit) (*entity.ConsumerLimit, error)
	Update(ctx context.Context, consumerLimit *entity.ConsumerLimit, updatedFields map[string]interface{}) (*entity.ConsumerLimit, error)
}

func (c *ConsumerLimitRepository) FindByConsumerId(ctx context.Context, consumerId uint64) ([]*entity.ConsumerLimit, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerLimitRepository - FindByConsumerId")
	defer span.End()

	var consumerLimits []*entity.ConsumerLimit
	if err := c.db.Debug().WithContext(ctxSpan).Where("consumer_id = ?", consumerId).Find(&consumerLimits).Error; err != nil {
		log.Println("ERROR: [ConsumerLimitRepository - FindByConsumerId] Internal server error:", err)
		return nil, err
	}

	return consumerLimits, nil
}

func (c *ConsumerLimitRepository) FindByConsumerIdAndTenor(ctx context.Context, consumerId uint64, tenor uint32) (*entity.ConsumerLimit, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerLimitRepository - FindByConsumerIdAndTenor")
	defer span.End()

	var consumerLimit entity.ConsumerLimit
	if err := c.db.Debug().WithContext(ctxSpan).Where("consumer_id = ? AND tenor = ?", consumerId, tenor).First(&consumerLimit).Error; err != nil {
		log.Println("ERROR: [ConsumerLimitRepository - FindByConsumerIdAndTenor] Internal server error:", err)
		return nil, err
	}

	return &consumerLimit, nil
}

func (c *ConsumerLimitRepository) Create(ctx context.Context, req *entity.ConsumerLimit) (*entity.ConsumerLimit, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerLimitRepository - Create")
	defer span.End()

	if err := c.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		log.Println("ERROR: [ConsumerLimitRepository - Create] Internal server error:", err)
		return nil, err
	}

	return req, nil
}

func (c *ConsumerLimitRepository) Update(ctx context.Context, consumerLimit *entity.ConsumerLimit, updatedFields map[string]interface{}) (*entity.ConsumerLimit, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerLimitRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()

	if err := c.db.Debug().WithContext(ctxSpan).Model(&consumerLimit).Updates(updatedFields).Error; err != nil {
		log.Println("ERROR: [ConsumerLimitRepository - Update] Internal server error:", err)
		return nil, err
	}

	return consumerLimit, nil
}
