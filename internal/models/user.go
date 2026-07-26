package models

import (
	"github.com/velocitykode/velocity/orm"
)

// User model represents a user in the application.
//
// It is also this application's auth model: internal/app/bootstrap.go hands
// it to ormauth, which maps these columns onto auth.Authenticatable. Swapping
// authentication onto a different model means pointing that call at a
// different type.
type User struct {
	orm.Model[User]
	Name     string `orm:"column:name;type:varchar(255);not_null" json:"name"`
	Email    string `orm:"column:email;type:varchar(255);unique;not_null" json:"email"`
	Password string `orm:"column:password;type:varchar(255);not_null" json:"-"`

	// RememberToken backs the remember-me cookie. It is a pointer because
	// the column is nullable for users who have never used remember-me, and
	// scanning SQL NULL into a plain string is a driver error.
	RememberToken *string `orm:"column:remember_token;type:varchar(255)" json:"-"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

// Guarded opts out of velocity deny-by-default mass assignment with an empty
// denylist (allow-all, no Fillable acronym-zeroing). Name a column here to keep
// map-based writes from ever reaching it.
//
// Auth needs remember_token writable: the remember-me token is persisted
// through the ORM's map-based update path.
func (User) Guarded() []string { return nil }

// GetAuthIdentifier returns the primary key, used as the session/JWT subject.
func (u *User) GetAuthIdentifier() interface{} { return u.ID }

// GetAuthPassword returns the stored password hash.
func (u *User) GetAuthPassword() string { return u.Password }

// GetRememberToken returns the remember-me token, or "" when unset.
func (u *User) GetRememberToken() string {
	if u.RememberToken == nil {
		return ""
	}
	return *u.RememberToken
}

// SetRememberToken sets the remember-me token.
func (u *User) SetRememberToken(token string) { u.RememberToken = &token }
