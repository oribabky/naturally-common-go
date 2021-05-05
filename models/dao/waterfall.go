package dao

import (
	"github.com/google/uuid"
	"github.com/oribabky/naturally-common-go/models/domain"
)

type Waterfall struct {
	Id                  string
	Name                string
	Country             string
	X                   float32
	Y                   float32
	IsConfirmedLocation bool
	TotalHeight         *float32
	LongestDrop         *float32
	ImageUrl            string
	InfoUrl             string
}

func ToDomain(self Waterfall) domain.Waterfall {
	waterfall := domain.NewWaterfall(
		uuid.MustParse(self.Id),
		self.Name,
		self.Country,
		self.X,
		self.Y,
		self.IsConfirmedLocation,
		self.TotalHeight,
		self.LongestDrop,
		self.ImageUrl,
		self.InfoUrl,
	)
	return waterfall
}
