package model

import "time"

type Drop struct {
	ID          string     `json:"id"`
	Ciphertext  string     `json:"ciphertext"`
	RevealAt    *time.Time `json:"reveal_at"`
	KnockTarget *int       `json:"knock_target"`
	KnockCount  *int       `json:"knock_count"`
	Burned      *bool      `json:"burned"`
	CreatedAt   *time.Time `json:"created_at"`
}
