package model

type WikipediaAnnotation interface {
	GetPageID() uint64
	GetTitle() string
	GetSpot() string
	GetStart() uint
	GetEnd() uint
	GetLinkProbability() float64
	GetRho() float64
}
