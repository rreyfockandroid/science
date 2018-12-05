package data

import "time"

type BikeIdx struct {
	Name string
	AddDate time.Time
	Body string
	Destination *Destination
}

type Destination struct {
	Fr uint8
	Dh uint8
	Xc uint8
	AllMontain uint8
}