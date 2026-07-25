package models

import (
	"github.com/velocitykode/velocity/orm"
)

// User model represents a user in the application
type User struct {
	orm.Model[User]
	Name     string `orm:"column:name;type:varchar(255);not_null" json:"name"`
	Email    string `orm:"column:email;type:varchar(255);unique;not_null" json:"email"`
	Password string `orm:"column:password;type:varchar(255);not_null" json:"-"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

// Guarded opts out of velocity deny-by-default mass assignment with an empty
// denylist (allow-all, no Fillable acronym-zeroing). Name a column here to keep
// map-based writes from ever reaching it.
func (User) Guarded() []string { return nil }
