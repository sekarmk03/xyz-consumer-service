package handler

import (
	"context"
	"log"
	"net/http"
	"xyz-consumer-service/common/config"
	commonErr "xyz-consumer-service/common/error"
	"xyz-consumer-service/modules/consumer_limit/entity"
	"xyz-consumer-service/modules/consumer_limit/service"
	"xyz-consumer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConsumerLimitHandler struct {
	pb.UnimplementedConsumerLimitServiceServer
	config           config.Config
	consumerLimitSvc service.ConsumerLimitServiceUseCase
}

func NewConsumerLimitHandler(config config.Config, consumerLimitSvc service.ConsumerLimitServiceUseCase) *ConsumerLimitHandler {
	return &ConsumerLimitHandler{
		config:           config,
		consumerLimitSvc: consumerLimitSvc,
	}
}

func (clh *ConsumerLimitHandler) GetConsumerLimitsByConsumerId(ctx context.Context, req *pb.ConsumerRequest) (*pb.ConsumerLimitListResponse, error) {
	consumerLimitList, err := clh.consumerLimitSvc.FindByConsumerId(ctx, req.ConsumerId)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitHandler - GetConsumerLimitsByConsumerId] Error while find consumer limits by consumer id:", parseError.Message)
		return &pb.ConsumerLimitListResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var consumerLimits []*pb.ConsumerLimit
	for _, cl := range consumerLimitList {
		consumerLimits = append(consumerLimits, entity.ConvertEntityToProto(cl))
	}

	return &pb.ConsumerLimitListResponse{
		Code:    uint32(http.StatusOK),
		Message: "Success get all consumer limits",
		Data:    consumerLimits,
	}, nil
}

func (clh *ConsumerLimitHandler) CreateConsumerLimit(ctx context.Context, req *pb.ConsumerLimit) (*pb.ConsumerLimitResponse, error) {
	consumerLimit, _ := clh.consumerLimitSvc.FindByConsumerIdAndTenor(ctx, req.ConsumerId, req.Tenor)
	if consumerLimit != nil {
		log.Println("WARNING: [ConsumerLimitHandler - CreateConsumerLimit] Consumer limit already exist for consumer id:", req.ConsumerId, "and tenor:", req.Tenor)
		return &pb.ConsumerLimitResponse{
			Code:    uint32(http.StatusConflict),
			Message: "Consumer limit already exist",
		}, status.Errorf(codes.AlreadyExists, "Consumer limit already exist")
	}

	consumerLimit, err := clh.consumerLimitSvc.Create(ctx, req.ConsumerId, req.Tenor, req.LimitAmount)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitHandler - CreateConsumerLimit] Error while create consumer limit:", parseError.Message)
		return &pb.ConsumerLimitResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.ConsumerLimitResponse{
		Code:    uint32(http.StatusCreated),
		Message: "Success create consumer limit",
		Data:    entity.ConvertEntityToProto(consumerLimit),
	}, nil
}

func (clh *ConsumerLimitHandler) UpdateAvailableLimit(ctx context.Context, req *pb.UpdateAvailableLimitRequest) (*pb.ConsumerLimitResponse, error) {
	consumerLimit, err := clh.consumerLimitSvc.FindByConsumerIdAndTenor(ctx, req.ConsumerId, req.Tenor)
	if err != nil {
		if consumerLimit == nil {
			log.Println("WARNING: [ConsumerLimitHandler - UpdateConsumerLimitAvailable] Consumer limit not found for consumer id:", req.ConsumerId, "and tenor:", req.Tenor)
			return &pb.ConsumerLimitResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "Consumer limit not found",
			}, status.Errorf(codes.NotFound, "Consumer limit not found")
		}
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitHandler - UpdateConsumerLimitAvailable] Error while find consumer limit by consumer id and tenor:", parseError.Message)
		return &pb.ConsumerLimitResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	consumerLimitUpdate, err := clh.consumerLimitSvc.UpdateLimitAvailable(ctx, consumerLimit, req.AmountTransaction)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerLimitHandler - UpdateConsumerLimitAvailable] Error while update consumer limit available:", parseError.Message)
		return &pb.ConsumerLimitResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.ConsumerLimitResponse{
		Code:    uint32(http.StatusOK),
		Message: "Success update consumer limit available",
		Data:    entity.ConvertEntityToProto(consumerLimitUpdate),
	}, nil
}
