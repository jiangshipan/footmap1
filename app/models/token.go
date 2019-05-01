package models

import "time"

type Token struct {
	Uid int
	Status int
	Tokens string
	Expired time.Time
}