package akin

import "time"

// Model base model to implement soft delete for Beego ORM
type Model struct {
	CreatedAt *time.Time `json:"created_at,omitempty" orm:"column(created_at);auto_now_add;type(timestamp with time zone)"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" orm:"column(updated_at);type(timestamp with time zone);null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" orm:"column(deleted_at);type(timestamp with time zone);;null"`
}
