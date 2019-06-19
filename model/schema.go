package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Company Schema
type Company struct {
	gorm.Model
	Name      string `gorm:"column:company_name"`
	Address   string `gorm:"column:address"`
	ContactNo string `gorm:"column:contact_no"`
}

// Users Schema
type Users struct {
	gorm.Model
	// Accessor for column field | Type of Column | column: x specifies the name for this field in the db
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	CompanyID uint   `gorm:"company_id"` // Foreign Key for company table
	// Specify related tables here.
	Company Company `gorm:"association_autoupdate:false;association_autocreate:false"`
}

// Use the below command to AutoMigrate structs to db tables
// db.AutoMigrate(&Company{})
