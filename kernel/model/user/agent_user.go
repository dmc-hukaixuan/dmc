package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int       `json:"user_id" gorm:"comment: agent ID"`
	Login      string    `json:"UserLogin" gorm:""`
	PW         string    `json:"pw" gorm:""`
	Title      string    `json:"title" gorm:""`
	Fullname   string    `json:"full_name" gorm:""`
	Email      string    `json:"email" gorm:""`
	Mobile     string    `json:"mobile" gorm:""`
	JobNumber  string    `json:"job_number" gorm:""`
	FirstName  string    `json:"first_name" gorm:""`
	LastName   string    `json:"last_name" gorm:""`
	DistrictID string    `json:"district_id" gorm:""`
	City       string    `json:"city" gorm:""`
	ValidID    string    `json:"valid_id" gorm:""`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;"` // 工单的创建时间
	CreateBy   int       `json:"create_by" gorm:"comment:api组"`      // 工单的创建人
	ChangeTime time.Time `json:"change_time" gorm:"autoUpdateTime;"` // 工单的修改时间
	ChangeBy   int       `json:"change_by" gorm:"comment:工单的修改人"`    // 工单的修改人
}

type LoginParam struct {
	Login    string `json:"login" gorm:"column:login;comment:工单的修改人"`
	PW       string `json:"pw" gorm:"column:pw;comment:工单的修改人"`
	UserType string `json:"pw" gorm:"-"`
}

type Role struct {
	ID                  int            `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name                string         `gorm:"column:name;NOT NULL" json:"name"`
	DefaultOwner        sql.NullInt32  `gorm:"column:default_owner" json:"default_owner"`
	DefaultResponsible  sql.NullInt32  `gorm:"column:default_responsible" json:"default_responsible"`
	UnlockTimeout       sql.NullInt32  `gorm:"column:unlock_timeout" json:"unlock_timeout"`
	FirstResponseTime   sql.NullInt32  `gorm:"column:first_response_time" json:"first_response_time"`
	FirstResponseNotify sql.NullInt32  `gorm:"column:first_response_notify" json:"first_response_notify"`
	UpdateTime          sql.NullInt32  `gorm:"column:update_time" json:"update_time"`
	UpdateNotify        sql.NullInt32  `gorm:"column:update_notify" json:"update_notify"`
	SolutionTime        sql.NullInt32  `gorm:"column:solution_time" json:"solution_time"`
	SolutionNotify      sql.NullInt32  `gorm:"column:solution_notify" json:"solution_notify"`
	SystemAddressId     sql.NullInt32  `gorm:"column:system_address_id" json:"system_address_id"`
	CalendarName        sql.NullString `gorm:"column:calendar_name" json:"calendar_name"`
	DefaultSignKey      sql.NullString `gorm:"column:default_sign_key" json:"default_sign_key"`
	SalutationId        sql.NullInt32  `gorm:"column:salutation_id" json:"salutation_id"`
	SignatureId         sql.NullInt32  `gorm:"column:signature_id" json:"signature_id"`
	FollowUpId          sql.NullInt32  `gorm:"column:follow_up_id" json:"follow_up_id"`
	FollowUpLock        sql.NullInt32  `gorm:"column:follow_up_lock" json:"follow_up_lock"`
	Comments            sql.NullString `gorm:"column:comments" json:"comments"`
	ValidId             int            `gorm:"column:valid_id;NOT NULL" json:"valid_id"`
	CreateTime          time.Time      `gorm:"column:create_time;NOT NULL" json:"create_time"`
	CreateBy            int            `gorm:"column:create_by;NOT NULL" json:"create_by"`
	ChangeTime          time.Time      `gorm:"column:change_time;NOT NULL" json:"change_time"`
	ChangeBy            int            `gorm:"column:change_by;NOT NULL" json:"change_by"`
}

func (m *Role) TableName() string {
	return "`Role`"
}
