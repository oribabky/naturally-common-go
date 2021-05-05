package domain

import "github.com/google/uuid"

type Waterfall interface {
	Id() uuid.UUID
	Name() string
	Country() string
	X() float32
	Y() float32
	TotalHeight() *float32
	LongestDrop() *float32
	ImageUrl() string
	InfoUrl() string
}

type waterfall struct {
	id                  uuid.UUID
	name                string
	country             string
	x                   float32
	y                   float32
	isConfirmedLocation bool
	totalHeight         *float32
	longestDrop         *float32
	imageUrl            string
	infoUrl             string
}

func NewWaterfall(
	id uuid.UUID,
	name string,
	country string,
	x float32,
	y float32,
	isConfirmedLocation bool,
	totalHeight *float32,
	longestDrop *float32,
	imageUrl string,
	InfoUrl string) Waterfall {
	return &waterfall{
		id:                  id,
		name:                name,
		country:             country,
		x:                   x,
		y:                   y,
		isConfirmedLocation: isConfirmedLocation,
		totalHeight:         totalHeight,
		longestDrop:         longestDrop,
		imageUrl:            imageUrl,
		infoUrl:             InfoUrl,
	}
}

func (w waterfall) Id() uuid.UUID { return w.id }

func (w waterfall) Name() string { return w.name }

func (w waterfall) Country() string { return w.country }

func (w waterfall) X() float32 { return w.x }

func (w waterfall) Y() float32 { return w.y }

func (w waterfall) IsConfirmedLocation() bool { return w.isConfirmedLocation }

func (w waterfall) TotalHeight() *float32 { return w.totalHeight }

func (w waterfall) LongestDrop() *float32 { return w.longestDrop }

func (w waterfall) ImageUrl() string { return w.imageUrl }

func (w waterfall) InfoUrl() string { return w.infoUrl }
