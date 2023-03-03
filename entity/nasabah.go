package entity

import (
	"tugas3rpl/common"
)

type Nasabah struct {
	common.Model
	Name     string     `json:"name" binding:"required"`
	Email    string     `json:"email" binding:"required"`
	Address  string     `json:"address" binding:"required"`
	NoTelp   []NoTelp   `json:"no_telp,omitempty"`
	Rekening []Rekening `json:"rekening,omitempty"`
}

func (Nasabah) TableName() string {
	return "nasabah"
}
