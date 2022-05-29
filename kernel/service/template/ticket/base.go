package ticket

import (
	model "dmc/kernel/model/ticket"
)

type TF interface {
	TemplateEditRender() model.FeildData
	SearchSQLGet()
	EditFieldRender()
	SearchFieldRender()
}

func TicketStartandField(fieldtype string) TF {
	// 这里应该获取系统配置中的配置，实现工单单号的生成

	switch fieldtype {
	case "user":
		return &OwnerBase{}
	case "priority":
		return &Priority{}
	case "richtext":
		return &Richtext{}
	case "service":
		return &Service{}
	case "sla":
		return &SLA{}
	case "state":
		return &TicketState{}
	case "source":
		return &TicketSource{}
	case "type":
		return &TicketType{}
	case "time":
		return &TimeBase{}
	case "title":
		return &Title{}
	case "role":
		return &Role{}

	// case "aliyun-oss":
	// 	return &AliyunOSS{}
	default:
		return &Title{}
	}
}
