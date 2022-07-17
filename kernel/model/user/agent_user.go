package user

import (
	"time"
)

type User struct {
	ID         int       `json:"user_id" gorm:"comment: agent ID"`
	Login      string    `json:"UserLogin" gorm:"login"`
	PW         string    `json:"pw" gorm:"pw"`
	Title      string    `json:"title" gorm:"title"`
	FullName   string    `json:"full_name" gorm:"full_name"`
	Email      string    `json:"email" gorm:"email"`
	Mobile     string    `json:"mobile" gorm:"mobile"`
	JobNumber  string    `json:"job_number" gorm:"job_number"`
	FirstName  string    `json:"first_name" gorm:"first_name"`
	LastName   string    `json:"last_name" gorm:"last_name"`
	DistrictID string    `json:"district_id" gorm:"district_id"`
	Department string    `json:"department" gorm:"-"`
	City       string    `json:"city" gorm:"city"`
	ValidID    string    `json:"valid_id" gorm:"valid_id"`
	CreateTime time.Time `json:"create_time" gorm:"create_time;"`   // 工单的创建时间
	CreateBy   int       `json:"create_by" gorm:"create_by:api组"`   // 工单的创建人
	ChangeTime time.Time `json:"change_time" gorm:"change_time;"`   // 工单的修改时间
	ChangeBy   int       `json:"change_by" gorm:"change_by:工单的修改人"` // 工单的修改人
}

type LoginParam struct {
	Login    string `json:"login" gorm:"column:login;comment:工单的修改人"`
	PW       string `json:"pw" gorm:"column:pw;comment:工单的修改人"`
	UserType string `json:"pw" gorm:"-"`
}

type Role struct {
	ID                     int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name                   string `gorm:"column:name;NOT NULL" json:"name"`
	DefaultOwner           int    `gorm:"column:default_owner;default:null" json:"default_owner"`
	DefaultOwnerName       string `gorm:"-" json:"default_owner_name"`
	DefaultResponsible     int    `gorm:"column:default_responsible;default:null" json:"default_responsible"`
	DefaultResponsibleName string `gorm:"-" json:"default_responsible_name"`
	UnlockTimeout          int    `gorm:"column:unlock_timeout;default:0" json:"unlock_timeout"`
	FirstResponseTime      int    `gorm:"column:first_response_time;default:null" json:"first_response_time"`
	FirstResponseNotify    int    `gorm:"column:first_response_notify;default:0" json:"first_response_notify"`
	UpdateTime             int    `gorm:"column:update_time;default:0" json:"update_time"`
	UpdateNotify           int    `gorm:"column:update_notify;default:0" json:"update_notify"`
	SolutionTime           int    `gorm:"column:solution_time;default:0" json:"solution_time"`
	SolutionNotify         int    `gorm:"column:solution_notify;default:0" json:"solution_notify"`
	SystemAddressId        int    `gorm:"column:system_address_id;default:null" json:"system_address_id,omitempty"`
	CalendarName           int    `gorm:"column:calendar_name;default:null" json:"calendar_name,omitempty"`
	DefaultSignKey         string `gorm:"column:default_sign_key;default:null" json:"default_sign_key"`
	SalutationId           int    `gorm:"column:salutation_id;default:null" json:"salutation_id,omitempty"`
	SignatureId            int    `gorm:"column:signature_id;default:null" json:"signature_id,omitempty"`
	FollowUpId             int    `gorm:"column:follow_up_id;default:null" json:"follow_up_id,omitempty"`
	FollowUpLock           int    `gorm:"column:follow_up_lock" json:"follow_up_lock,omitempty"`
	Comments               string `gorm:"column:comments" json:"comments"`
	ValidID                int    `gorm:"column:valid_id;NOT NULL" json:"validid"`
	CreateTime             string `gorm:"column:create_time;NOT NULL" json:"create_time"`
	CreateBy               int    `gorm:"column:create_by;NOT NULL" json:"create_by"`
	CreateByName           string `gorm:"-" json:"create_by_name"`
	ChangeTime             string `gorm:"column:change_time;NOT NULL" json:"change_time,omitempty"`
	ChangeBy               int    `gorm:"column:change_by;NOT NULL" json:"change_by,omitempty"`
	ChangeByName           string `gorm:"-" json:"change_by_name,omitempty"`
}

type UserRole struct {
	ID           int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	QueueID      int    `gorm:"column:queue_id;default:null" json:"queue_id"`
	UserID       int    `gorm:"column:user_id;default:null" json:"user_id"`
	CreateTime   string `gorm:"column:create_time;NOT NULL" json:"create_time"`
	CreateBy     int    `gorm:"column:create_by;NOT NULL" json:"create_by"`
	CreateByName string `gorm:"-" json:"create_by_name"`
	ChangeTime   string `gorm:"column:change_time;NOT NULL" json:"change_time,omitempty"`
	ChangeBy     int    `gorm:"column:change_by;NOT NULL" json:"change_by,omitempty"`
	ChangeByName string `gorm:"-" json:"change_by_name,omitempty"`
}

type RoleTemplate struct {
	ID         int `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TemplateID int `gorm:"column:template_id;NOT NULL" json:"template_id"`
	RoleID     int `gorm:"column:role_id;NOT NULL" json:"role_id"`
}

func (m *Role) TableName() string {
	return "`Role`"
}
