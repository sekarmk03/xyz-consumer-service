package entity

import (
	"time"
	"xyz-consumer-service/pb"
)

const (
	ConsumerLimitTableName = "consumer_limits"
)

type ConsumerLimit struct {
	Id             uint64    `json:"id"`
	ConsumerId     uint64    `json:"consumer_id"`
	Tenor          uint64    `json:"tenor"`
	LimitAmount    uint64    `json:"limit_amount"`
	LimitAvailable uint64    `json:"limit_available"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (c *ConsumerLimit) TableName() string {
	return ConsumerLimitTableName
}

func ConvertEntityToProto(c *ConsumerLimit) *pb.ConsumerLimit {
	return &pb.ConsumerLimit{
		Id:             c.Id,
		ConsumerId:     c.ConsumerId,
		Tenor:          c.Tenor,
		LimitAmount:    c.LimitAmount,
		LimitAvailable: c.LimitAvailable,
		CreatedAt:      c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      c.UpdatedAt.Format(time.RFC3339),
	}
}
