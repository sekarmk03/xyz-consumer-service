package repository

import (
	"context"
	"errors"
	"log"
	"time"
	"xyz-consumer-service/modules/consumer/entity"

	"github.com/go-sql-driver/mysql"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ConsumerRepository struct {
	db *gorm.DB
}

func NewConsumerRepository(db *gorm.DB) *ConsumerRepository {
	return &ConsumerRepository{
		db: db,
	}
}

type ConsumerRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Consumer, error)
	FindById(ctx context.Context, id uint64) (*entity.Consumer, error)
	Create(ctx context.Context, req *entity.Consumer) (*entity.Consumer, error)
	Update(ctx context.Context, consumer *entity.Consumer, updatedFields map[string]interface{}) (*entity.Consumer, error)
	Delete(ctx context.Context, id uint64) error
}

func (c *ConsumerRepository) FindAll(ctx context.Context, req any) ([]*entity.Consumer, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerRepository - FindAll")
	defer span.End()

	var consumers []*entity.Consumer
	if err := c.db.Debug().WithContext(ctxSpan).Order("created_at desc").Find(&consumers).Error; err != nil {
		log.Println("ERROR: [ConsumerRepository - FindAll] Internal server error:", err)
		return nil, err
	}

	return consumers, nil
}

func (c *ConsumerRepository) FindById(ctx context.Context, id uint64) (*entity.Consumer, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerRepository - FindById")
	defer span.End()

	var consumer entity.Consumer
	if err := c.db.Debug().WithContext(ctxSpan).Where("id = ?", id).First(&consumer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [ConsumerRepository - FindById] Consumer not found for id:", id)
			return nil, status.Errorf(codes.NotFound, "Consumer not found for id: %v", id)
		}
		log.Println("ERROR: [ConsumerRepository - FindById] Internal server error:", err)
		return nil, err
	}

	return &consumer, nil
}

func (c *ConsumerRepository) Create(ctx context.Context, req *entity.Consumer) (*entity.Consumer, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerRepository - Create")
	defer span.End()

	if err := c.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			log.Println("WARNING: [ConsumerRepository - Create] Consumer already exists for nik:", req.Nik)
			return nil, status.Errorf(codes.AlreadyExists, "Consumer already exists for nik: %v", req.Nik)
		}
		log.Println("ERROR: [ConsumerRepository - Create] Internal server error:", err)
		return nil, err
	}

	return req, nil
}

func (c *ConsumerRepository) Update(ctx context.Context, consumer *entity.Consumer, updatedFields map[string]interface{}) (*entity.Consumer, error) {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()

	if err := c.db.Debug().WithContext(ctxSpan).Model(&consumer).Updates(updatedFields).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [ConsumerRepository - Update] Consumer not found for id:", consumer.Id)
			return nil, status.Errorf(codes.NotFound, "Consumer not found for id: %v", consumer.Id)
		}
		log.Println("ERROR: [ConsumerRepository - Update] Internal server error:", err)
		return nil, err
	}

	return consumer, nil
}

func (c *ConsumerRepository) Delete(ctx context.Context, id uint64) error {
	ctxSpan, span := trace.StartSpan(ctx, "ConsumerRepository - Delete")
	defer span.End()

	if err := c.db.Debug().WithContext(ctxSpan).Where("id = ?", id).Delete(&entity.Consumer{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [ConsumerRepository - Delete] Consumer not found for id:", id)
			return status.Errorf(codes.NotFound, "Consumer not found for id: %v", id)
		}
		log.Println("ERROR: [ConsumerRepository - Delete] Internal server error:", err)
		return err
	}

	return nil
}
