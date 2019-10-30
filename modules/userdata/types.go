package userdata

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Handler struct
type Handler struct {
	DB *gorm.DB
}

// User struct
type User struct {
	UserID    string
	Name      string `gorm:"type:varchar(128)"`
	MSISDN    string
	Email     string `gorm:"type:varchar(512)"`
	BirthDate time.Time
	UserAge   int
}
