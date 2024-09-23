package entity

import (
	"time"
	"xyz-consumer-service/pb"
)

const (
	ConsumerTableName = "consumers"
)

type Consumer struct {
	Id          uint64    `json:"id"`
	Nik         string    `json:"nik"`
	Fullname    string    `json:"fullname"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   string    `json:"birth_date"`
	Salary      uint64    `json:"salary"`
	Ktp         string    `json:"ktp"`
	SelfiePhoto string    `json:"selfie_photo"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *Consumer) TableName() string {
	return ConsumerTableName
}

func ConvertEntityToProto(c *Consumer) *pb.Consumer {
	return &pb.Consumer{
		Id:          c.Id,
		Nik:         c.Nik,
		Fullname:    c.Fullname,
		LegalName:   c.LegalName,
		BirthPlace:  c.BirthPlace,
		BirthDate:   c.BirthDate,
		Salary:      c.Salary,
		Ktp:         c.Ktp,
		SelfiePhoto: c.SelfiePhoto,
		CreatedAt:   c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   c.UpdatedAt.Format(time.RFC3339),
	}
}
