package entity

import "tugas3rpl/common"

type NoTelp struct {
	common.Model
	CountryCode string   `json:"country_code" binding:"required"`
	Number      string   `json:"number" binding:"required"`
	NasabahID   uint64   `gorm:"foreignKey" json:"nasabah_id" binding:"required"`
	Nasabah     *Nasabah `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"nasabah,omitempty"`
}

func (NoTelp) TableName() string {
	return "no_telp"
}
