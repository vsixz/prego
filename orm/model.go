package orm

import (
	"github.com/vsixz/prego/util"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PureBaseModel struct {
	Id string `gorm:"type.go:char(36);primary_key" json:"id"`
}

type BaseModel struct {
	Id        string    `gorm:"type.go:char(36);primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type IntegerBaseModel struct {
	Id        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	base.Id = util.SweepString(uuid.New().String(), "-")
	return nil
}
