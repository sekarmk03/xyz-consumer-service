package consumer

import (
	"xyz-consumer-service/common/config"
	"xyz-consumer-service/modules/consumer/internal/builder"
	"xyz-consumer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	consumer := builder.BuildConsumerHandler(cfg, db, grpcConn)
	pb.RegisterConsumerServiceServer(server, consumer)
}
