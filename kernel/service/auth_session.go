package service

import (
	"dmc/initialize/database"
	"dmc/kernel/model/user"
	"dmc/kernel/util"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type session struct {
	ID         int64  `json:"id" gorm:"column:id;comment:"`
	SessionID  string `json:"session_id" gorm:"column:session_id;comment:"`
	DataKey    string `json:"data_key" gorm:"column:data_key;comment:"`
	DataValue  string `json:"data_value" gorm:"column:data_value;comment:"`
	Serialized int    `json:"Serialized" gorm:"column:serialized;comment:"`
}

// 创建 seesion
func CreateSessionID(user *user.User) (token string, err error) {
	// time now
	n_time := time.Now()
	n_time_unix := strconv.Itoa(int(n_time.Unix()))
	n_time_stamp := n_time.Format("2006-01-02 15:04:09")
	// get remote address and the http user agent

	// create challenge token
	challengeToken := util.GenerateRandomString(32)

	// create session id
	sessionID := util.GenerateRandomString(32)
	var session = []session{
		{SessionID: sessionID, DataKey: "UserType", DataValue: "jinzhu3"},
		{SessionID: sessionID, DataKey: "UserSessionStart", DataValue: n_time_unix},
		{SessionID: sessionID, DataKey: "UserRemoteAddr", DataValue: "jinzhu3"},
		{SessionID: sessionID, DataKey: "UserRemoteUserAgent", DataValue: "jinzhu3"},
		{SessionID: sessionID, DataKey: "UserChallengeToken", DataValue: challengeToken},
		{SessionID: sessionID, DataKey: "UserLastRequest", DataValue: n_time_unix},
		{SessionID: sessionID, DataKey: "UserLastLogin", DataValue: n_time_unix},
		{SessionID: sessionID, DataKey: "UserLastLoginTimestamp", DataValue: n_time_stamp},
		{SessionID: sessionID, DataKey: "UserFullname", DataValue: user.Fullname},
		{SessionID: sessionID, DataKey: "UserLogin", DataValue: user.Login},
	}
	err = database.Gorm().Table("sessions").Create(&session).Error

	return sessionID, err
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

// 检查 session 是否过期，
// 这里应该把 session 写入缓存中，
//
func CheckSessionID(tokenString string) (user1 *user.User, err error) {

	sessionData, err := GetSessionIDData(tokenString)
	// 没有获取到数据
	if err != nil {
		return user1, err
	}
	var userInfo user.User
	time1, ok := sessionData["UserLastLoginTimestamp"]
	if ok {
		userInfo.Login = sessionData["UserLogin"]
		fmt.Println("time1 :", time1)
	} else {
		userInfo.Login = sessionData["UserLogin"]

	}
	fmt.Println("time1 :", time1)
	return &userInfo, err
}

// 获取 数据库中的 seesion 信息
func GetSessionIDData(tokenString string) (sessionData map[string]string, err error) {
	var session []session
	// ask database to session data
	database.Gorm().Table("sessions").Where("session_id = ? ", tokenString).Find(&session)

	datasession := make(map[string]string)
	// do loop get detail session data info
	for _, v := range session {
		datasession[v.DataKey] = v.DataValue
	}

	return datasession, err
}
