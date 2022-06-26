package number

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type TN interface {
	TicketNumberBuild() string
}

// @author:
// @param:
// @description:
func TicketNumber() TN {
	// 这里应该获取系统配置中的配置，实现工单单号的生成
	NumberGenerator := "dateChecksum"
	switch NumberGenerator {
	case "date":
		return &Date{}
	case "autoIncrement":
		return &AutoIncrement{}
	case "dateChecksum":
		return &DateChecksum{}
	// case "aliyun-oss":
	// 	return &AliyunOSS{}
	default:
		return &Date{}
	}
}
