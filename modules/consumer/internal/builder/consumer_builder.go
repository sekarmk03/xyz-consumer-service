package builder

import (
	"xyz-consumer-service/common/config"
	"xyz-consumer-service/modules/consumer/internal/handler"
	"xyz-consumer-service/modules/consumer/internal/repository"
	"xyz-consumer-service/modules/consumer/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildConsumerHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.ConsumerHandler {
	consumerRepository := repository.NewConsumerRepository(db)
	imageSvc := service.NewImageService(cfg)
	consumerSvc := service.NewConsumerService(cfg, consumerRepository)

	return handler.NewConsumerHandler(cfg, consumerSvc, imageSvc)
}
