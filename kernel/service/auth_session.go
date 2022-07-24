package service

import (
    "dmc/global"
    "dmc/global/log"
    "dmc/kernel/model/common/request"
    "dmc/kernel/model/user"
    "dmc/kernel/util"
    dateTime "dmc/kernel/util/time"
    "encoding/json"
    "errors"
    "fmt"
    "time"
)

// 创建 seesion
func CreateSessionID(user *user.User, RemoteAddr string, RemoteUserAgent string) (token string, err error) {
    // time now
    // n_time := dateTime.CurrentTimestamp()
    // //n_time_unix := strconv.Itoa(int(n_time.Unix()))
    // n_time_stamp := dateTime.CurrentTimestamp()

    // get remote address and the http user agent

    // create challenge token
    challengeToken := util.GenerateRandomString(32)

    userstr, _ := json.Marshal(user)

    // create session id
    sessionID := util.GenerateRandomString(32)
    session := request.Session{
        UserID:          user.ID,
        SessionID:       sessionID,
        ChallengeToken:  challengeToken,
        LoginTime:       time.Now(),
        ExpiresTime:     time.Now(),
        LastRequest:     time.Now(),
        RemoteAddr:      RemoteAddr,
        RemoteUserAgent: RemoteUserAgent,
        DataKey:         "",
        DataValue:       string(userstr),
        Serialized:      1,
    }
    fmt.Println("-----------: ", user.ID, user)
    // var session = []session{
    // 	{SessionID: sessionID, DataKey: "UserType", DataValue: "jinzhu3"},
    // 	{SessionID: sessionID, DataKey: "UserSessionStart", DataValue: n_time_unix},
    // 	{SessionID: sessionID, DataKey: "UserRemoteAddr", DataValue: "jinzhu3"},
    // 	{SessionID: sessionID, DataKey: "UserRemoteUserAgent", DataValue: "jinzhu3"},
    // 	{SessionID: sessionID, DataKey: "UserChallengeToken", DataValue: challengeToken},
    // 	{SessionID: sessionID, DataKey: "UserLastRequest", DataValue: n_time_unix},
    // 	{SessionID: sessionID, DataKey: "UserLastLogin", DataValue: n_time_unix},
    // 	{SessionID: sessionID, DataKey: "UserLastLoginTimestamp", DataValue: n_time_stamp},
    // 	{SessionID: sessionID, DataKey: "UserFullname", DataValue: user.FullName},
    // 	{SessionID: sessionID, DataKey: "UserLogin", DataValue: user.Login},
    // }

    err = global.GVA_DB.Table("sessions").Create(&session).Error
    fmt.Println("err", err)
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
func CheckSessionID(tokenString string, RemoteAddr string, RemoteUserAgent string) (user1 int, message string, sessionData request.Session) {
    // get session data
    sessionData = GetSessionIDData(tokenString)
    fmt.Println("tokenString ********   ", tokenString)
    // remote ip check
    // $ConfigObject->Get('SessionCheckRemoteIP')
    if sessionData.RemoteAddr != RemoteAddr {
        log.SugarLogger.Infof("RemoteIP of %s (%s) is different from registered IP (%s). Invalidating session!  Disable config 'SessionCheckRemoteIP' if you don't want this!", tokenString, sessionData.RemoteAddr, RemoteAddr)
        // delete session id if it isn't the same remote ip?
        return
    }
    nowTime := dateTime.NowSystemTime()

    MaxSessionIdleTime := 1000000
    SessionIdleTime := time.Hour * time.Duration(MaxSessionIdleTime)
    // check session idle time
    if sessionData.LastRequest.Add(SessionIdleTime).Before(nowTime) {
        message = "Session has timed out. Please log in again."

        timeOut := nowTime.Sub(sessionData.LastRequest)
        log.SugarLogger.Infof("SessionID (%s) idle timeout (%s)! Don't grant access!!!", tokenString, timeOut)
        // delete session id if too old
        // RemoveSessionID
        // $Self->RemoveSessionID( SessionID => $Param{SessionID} );
        return 0, message, sessionData
    }
    // check session time
    SessionMaxTime := 400000
    //var Hour time.Duration
    Hour := time.Hour * time.Duration(SessionMaxTime)
    if sessionData.LoginTime.Add(Hour).Before(nowTime) {
        message = "Session has timed out. Please log in again."
        timeOut := nowTime.Sub(sessionData.LoginTime)
        log.SugarLogger.Infof("SessionID (%s) idle timeout (%s)! Don't grant access!!!", tokenString, timeOut)
        // delete session id if too old
        // RemoveSessionID
        // $Self->RemoveSessionID( SessionID => $Param{SessionID} );
        return 0, message, sessionData
    }

    return 1, "", sessionData
}

// 获取 数据库中的 seesion 信息
func GetSessionIDData(tokenString string) (sessionData request.Session) {

    // ask database to session data
    err := global.GVA_DB.Table("sessions").Where("session_id = ? ", tokenString).Find(&sessionData).Error
    if err != nil {
        panic(err)
    }
    // datasession := make(map[string]string)
    // // do loop get detail session data info
    // for _, v := range session {
    // 	datasession[v.DataKey] = v.DataValue
    // }

    return sessionData
}
