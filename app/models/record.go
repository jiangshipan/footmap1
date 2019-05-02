package models

import "time"

type Record struct {
	Rid int 			`json:"rid"`
	Uid int				`json:"uid"`
	StartTime time.Time `json:"start_time"`
	TargetTime time.Time`json:"target_time"`
	Remarks string 		`json:"remarks"`
	TargetPlace string	`json:"target_place"`
	Status int			`json:"status"`
	Title string		`json:"title"`
}