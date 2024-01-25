package orm

import (
	"gorm.io/gorm"
)

// Model Tbl_User
type Tbl_User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}
