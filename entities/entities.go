package entities

import "time"

type Juice struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Expiration time.Time `json:"expiration"`
}
