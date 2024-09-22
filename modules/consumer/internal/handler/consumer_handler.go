package handler

import (
	"context"
	"log"
	"net/http"
	"xyz-consumer-service/common/config"
	commonErr "xyz-consumer-service/common/error"
	"xyz-consumer-service/modules/consumer/entity"
	"xyz-consumer-service/modules/consumer/service"
	"xyz-consumer-service/pb"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConsumerHandler struct {
	pb.UnimplementedConsumerServiceServer
	config      config.Config
	consumerSvc service.ConsumerServiceUseCase
	imageSvc    service.ImageServiceUseCase
}

func NewConsumerHandler(config config.Config, consumerSvc service.ConsumerServiceUseCase, imageSvc service.ImageServiceUseCase) *ConsumerHandler {
	return &ConsumerHandler{
		config:      config,
		consumerSvc: consumerSvc,
		imageSvc:    imageSvc,
	}
}

func (ch *ConsumerHandler) GetAllConsumers(ctx context.Context, req *emptypb.Empty) (*pb.ConsumerListResponse, error) {
	consumerList, err := ch.consumerSvc.FindAll(ctx, req)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - GetAllConsumers] Error while find all consumer:", parseError.Message)
		return &pb.ConsumerListResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var consumers []*pb.Consumer
	for _, c := range consumerList {
		consumers = append(consumers, entity.ConvertEntityToProto(c))
	}

	return &pb.ConsumerListResponse{
		Code:      uint32(http.StatusOK),
		Message:   "Success get all consumers",
		Consumers: consumers,
	}, nil
}

func (ch *ConsumerHandler) GetConsumerById(ctx context.Context, req *pb.ConsumerIdRequest) (*pb.ConsumerResponse, error) {
	consumer, err := ch.consumerSvc.FindById(ctx, req.Id)
	if err != nil {
		if consumer == nil {
			log.Println("WARNING: [ConsumerHandler - GetConsumerById] Consumer not found for id:", req.Id)
			return &pb.ConsumerResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "Consumer not found",
			}, status.Errorf(http.StatusNotFound, "Consumer not found for id: %v", req.Id)
		}
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - GetConsumerById] Internal server error:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.ConsumerResponse{
		Code:     uint32(http.StatusOK),
		Message:  "Success get consumer by id",
		Consumer: entity.ConvertEntityToProto(consumer),
	}, nil
}

func (ch *ConsumerHandler) CreateConsumer(ctx context.Context, req *pb.ConsumerDataRequest) (*pb.ConsumerResponse, error) {
	ktp, err := ch.imageSvc.UploadImage(ctx, req.KtpFilename, req.KtpBuffer)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - CreateConsumer] Error while upload ktp image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	selfiePhoto, err := ch.imageSvc.UploadImage(ctx, req.SelfiePhotoFilename, req.SelfiePhotoBuffer)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - CreateConsumer] Error while upload selfie photo image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	consumer, err := ch.consumerSvc.Create(ctx, req.Nik, req.Fullname, req.LegalName, req.BirthPlace, req.BirthDate, req.Salary, ktp, selfiePhoto)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - CreateConsumer] Error while create consumer:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.ConsumerResponse{
		Code:     uint32(http.StatusCreated),
		Message:  "Success create consumer",
		Consumer: entity.ConvertEntityToProto(consumer),
	}, nil
}

func (ch *ConsumerHandler) UpdateConsumer(ctx context.Context, req *pb.ConsumerDataRequest) (*pb.ConsumerResponse, error) {
	consumer, err := ch.consumerSvc.FindById(ctx, req.Id)
	if err != nil {
		if consumer == nil {
			log.Println("WARNING: [ConsumerHandler - UpdateConsumer] Consumer not found for id:", req.Id)
			return &pb.ConsumerResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "Consumer not found",
			}, status.Errorf(http.StatusNotFound, "Consumer not found for id: %v", req.Id)
		}
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - UpdateConsumer] Error while find consumer by id:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	ktp, err := ch.imageSvc.UploadImage(ctx, req.KtpFilename, req.KtpBuffer)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - UpdateConsumer] Error while upload ktp image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ch.imageSvc.DeleteImage(ctx, consumer.Ktp)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - UpdateConsumer] Error while delete ktp image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	selfiePhoto, err := ch.imageSvc.UploadImage(ctx, req.SelfiePhotoFilename, req.SelfiePhotoBuffer)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - UpdateConsumer] Error while upload selfie photo image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ch.imageSvc.DeleteImage(ctx, consumer.SelfiePhoto)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - UpdateConsumer] Error while delete selfie photo image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	consumer.Nik = req.Nik
	consumer.Fullname = req.Fullname
	consumer.LegalName = req.LegalName
	consumer.BirthPlace = req.BirthPlace
	consumer.BirthDate = req.BirthDate
	consumer.Salary = req.Salary
	consumer.Ktp = ktp
	consumer.SelfiePhoto = selfiePhoto

	res, err := ch.consumerSvc.Update(ctx, req.Id, consumer)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - UpdateConsumer] Error while update consumer:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.ConsumerResponse{
		Code:     uint32(http.StatusOK),
		Message:  "Success update consumer",
		Consumer: entity.ConvertEntityToProto(res),
	}, nil
}

func (ch *ConsumerHandler) DeleteConsumer(ctx context.Context, req *pb.ConsumerIdRequest) (*pb.ConsumerResponse, error) {
	consumer, err := ch.consumerSvc.FindById(ctx, req.Id)
	if err != nil {
		if consumer == nil {
			log.Println("WARNING: [ConsumerHandler - DeleteConsumer] Consumer not found for id:", req.Id)
			return &pb.ConsumerResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "Consumer not found",
			}, status.Errorf(http.StatusNotFound, "Consumer not found for id: %v", req.Id)
		}
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - DeleteConsumer] Error while find consumer by id:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ch.imageSvc.DeleteImage(ctx, consumer.Ktp)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - DeleteConsumer] Error while delete ktp image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ch.imageSvc.DeleteImage(ctx, consumer.SelfiePhoto)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - DeleteConsumer] Error while delete selfie photo image:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ch.consumerSvc.Delete(ctx, req.Id)
	if err != nil {
		parseError := commonErr.ParseError(err)
		log.Println("ERROR: [ConsumerHandler - DeleteConsumer] Error while delete consumer:", parseError.Message)
		return &pb.ConsumerResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.ConsumerResponse{
		Code:    uint32(http.StatusOK),
		Message: "Success delete consumer",
	}, nil
}
