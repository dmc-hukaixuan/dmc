package request

import (
	"time"

	"github.com/google/uuid"
)

// Custom claims structure
type TokenUserInfo struct {
	BaseTokenInfo
	BufferTime int64
}

type BaseTokenInfo struct {
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
}

type SubActionData struct {
	SubAction string                 `json:"subaction"`
	Data      map[string]interface{} `json:"data"`
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	//jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
}

type SessionData struct {
	UserID    int
	FirstName string
	LastName  string
	FullName  string
	Roles     []int `json:"roles"`
}

type Session struct {
	ID              int64     `json:"id" gorm:"column:id;comment:"`
	UserID          int       `json:"user_id" gorm:"column:user_id;comment:"`
	SessionID       string    `json:"session_id" gorm:"column:session_id;comment:"`
	ChallengeToken  string    `json:"challenge_token" gorm:"column:challenge_token;comment:"`
	LoginTime       time.Time `json:"login_time" gorm:"column:login_time;comment:"`
	ExpiresTime     time.Time `json:"expires_time" gorm:"column:expires_time;comment:"`
	LastRequest     time.Time `json:"last_request" gorm:"column:last_request;comment:"`
	RemoteAddr      string    `json:"remote_addr" gorm:"column:remote_addr;comment:"`
	RemoteUserAgent string    `json:"remote_user_agent" gorm:"column:remote_user_agent;comment:"`
	DataKey         string    `json:"data_key" gorm:"column:data_key;comment:"`
	DataValue       string    `json:"data_value" gorm:"column:data_value;comment:"`
	Serialized      int       `json:"Serialized" gorm:"column:serialized;comment:"`
}
