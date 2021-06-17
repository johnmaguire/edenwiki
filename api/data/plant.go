package data

import "time"

type PlantType string

const (
	TomatoPlant     PlantType = "Tomato"
	ZucchiniPlant             = "Zucchini"
	StrawberryPlant           = "Strawberry"
	PepperPlant               = "Pepper"
	CarrotPlant               = "Carrot"
	RadishPlant               = "Radish"
	CeleryPlant               = "Celery"
)

type Plot struct {
	Width  int
	Length int
	Plants map[Point]Planting
}

type Point struct {
	X int
	Y int
}

type Planting struct {
	Plants []Plant
}

type Plant struct {
	Type    PlantType
	Variety string
	Notes   []Note
}

type Note struct {
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
