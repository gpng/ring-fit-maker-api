package usersvc

import (
	vr "github.com/gpng/ring-fit-maker/utils/validator"
	"github.com/jinzhu/gorm"
)

// Service struct
type Service struct {
	db        *gorm.DB
	validator *vr.Validator
}

// New service
func New(
	db *gorm.DB,
	validator *vr.Validator,
) *Service {
	return &Service{
		db,
		validator,
	}
}
