package request

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
