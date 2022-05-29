package user

import "time"

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
