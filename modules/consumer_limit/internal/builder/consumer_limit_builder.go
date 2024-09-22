package builder

import (
	"xyz-consumer-service/common/config"
	"xyz-consumer-service/modules/consumer_limit/internal/handler"
	"xyz-consumer-service/modules/consumer_limit/internal/repository"
	"xyz-consumer-service/modules/consumer_limit/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildConsumerLimitHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.ConsumerLimitHandler {
	consumerLimitRepository := repository.NewConsumerLimitRepository(db)
	consumerLimitSvc := service.NewConsumerLimitService(cfg, consumerLimitRepository)

	return handler.NewConsumerLimitHandler(cfg, consumerLimitSvc)
}
