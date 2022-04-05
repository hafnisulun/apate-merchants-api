package models

import (
	"os"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Merchant struct {
	Base
	Name          string    `json:"name" gorm:"index;not null;"`
	Lat           float64   `json:"lat" gorm:"not null;"`
	Lon           float64   `json:"lon" gorm:"not null;"`
	Phone         string    `json:"phone"`
	Image         string    `json:"image"`
	Address       string    `json:"address"`
	ResidenceUUID uuid.UUID `json:"residence_uuid" gorm:"index;not null"`
	ClusterUUID   uuid.UUID `json:"cluster_uuid" gorm:"index;not null"`
}

type CreateMerchantInput struct {
	Name          string    `json:"name" binding:"required"`
	Lat           float64   `json:"lat" binding:"required"`
	Lon           float64   `json:"lon" binding:"required"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	ResidenceUUID uuid.UUID `json:"residence_uuid" binding:"required"`
	ClusterUUID   uuid.UUID `json:"cluster_uuid" binding:"required"`
}

type FindMerchantsQuery struct {
	ResidenceUUID string `form:"residence_uuid"`
	Page          int    `form:"page"`
	PerPage       int    `form:"per_page"`
}

func (merchant *Merchant) AfterFind(tx *gorm.DB) (err error) {
	if merchant.Image != "" {
		merchant.Image = os.Getenv("APATE_STORAGE_BASE_URL") + "/static/merchants/" + merchant.Image
	}
	return
}
