package consumerlimit

import (
	"xyz-consumer-service/common/config"
	"xyz-consumer-service/modules/consumer_limit/internal/builder"
	"xyz-consumer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	consumerLimit := builder.BuildConsumerLimitHandler(cfg, db, grpcConn)
	pb.RegisterConsumerLimitServiceServer(server, consumerLimit)
}
