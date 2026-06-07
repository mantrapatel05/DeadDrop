package model

import (
	"errors"
	"time"
)

type Drop struct {
	ID          string     `json:"id"`
	Ciphertext  string     `json:"-"`
	Nonce       string     `json:"-"`
	RevealAt    *time.Time `json:"reveal_at"`
	KnockTarget *int       `json:"knock_target"`
	KnockCount  *int       `json:"knock_count"`
	Burned      *bool      `json:"burned"`
	CreatedAt   *time.Time `json:"created_at"`
}
//3 funcs for drop 
//drop ready to use
//drop already  burned
//valid data for creqating a new drop

func (d *Drop) IsBurned() bool{
	if d.Burned == nil {
		return false
	}
	return *d.Burned
}

func (d *Drop) IsReady() bool {
	//burn drop not ready
	if d.IsBurned() {
		return false
	}

	// time unlock
	if d.RevealAt != nil && time.Now().After(*d.RevealAt) {
		return true
	}

	// knock unlock
	if d.KnockTarget != nil && d.KnockCount != nil {
		if *d.KnockCount >= *d.KnockTarget {
			return true
		}
	}

	return false
}

func (d *Drop) ValidCreate() error {
	if d.KnockTarget == nil && d.RevealAt == nil {
		return errors.New("drop must have at least one unlock method")
	}
	if d.KnockTarget != nil && *d.KnockTarget <= 0 {
		return errors.New("knock target must be greater than 0")
	}

	return nil
}