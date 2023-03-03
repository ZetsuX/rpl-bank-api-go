package entity

import "tugas3rpl/common"

type Rekening struct {
	common.Model
	Balance   float64  `json:"balance" binding:"required" sql:"type:numeric(19,2);"`
	Type      string   `json:"type" binding:"required"`
	Active    bool     `json:"active" binding:"required"`
	NasabahID uint64   `gorm:"foreignKey" json:"nasabah_id" binding:"required"`
	Nasabah   *Nasabah `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"nasabah,omitempty"`
}

func (Rekening) TableName() string {
	return "rekening"
}
