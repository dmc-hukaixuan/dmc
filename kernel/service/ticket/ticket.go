package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	number "dmc/kernel/system/ticket/number"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
	创建工单函数
*/
func TicketCreate(TicketBaseData model.TicketBaseData) (TicketID int64, err error) {
	// get db object
	// db := pool.GetDB()
	numberbuild := number.TicketNumber()
	// check ticket number if not null, generate ticket number
	if TicketBaseData.TicketNumber == "" {
		TicketBaseData.TicketNumber = numberbuild.TicketNumberBuild()
	}
	fmt.Println("TicketBaseData ", TicketBaseData)

	if !errors.Is(global.GVA_DB.Table("sc_ticket").Create(&TicketBaseData).Error, gorm.ErrRecordNotFound) {
		return 0, errors.New("Ticket create failed")
	}
	// update dynamic_field

	// trgger event loop
	// trigger event

	return TicketBaseData.ID, err
}

/*
	获取g单所有动态字段的值
*/
func TicketynamicFieldValueGet(ticketid int64) map[string]string {
	type DynamicFieldValue struct {
		Label     string `json:"label"`
		Name      string `json:"name"`
		ValueText string `json:"value_text"`
	}
	var dfv []DynamicFieldValue
	//db := pool.GetDB()
	// sql string TicketHistoryGet ticket_history_2022_02
	SelctSQL := `SELECT dfv.id, dfv.field_id, df.name as name, df.label as label, max(case when dfv.value_text IS NOT NULL then dfv.value_text ELSE dfv.value_date END) as value_text
				FROM dynamic_field_value dfv LEFT JOIN dynamic_field df ON df.id = dfv.field_id WHERE dfv.object_id = ?
				GROUP BY dfv.id`
	// ask database and fetch result
	global.GVA_DB.Raw(SelctSQL, ticketid).Scan(&dfv)
	dynamic_field := make(map[string]string)
	for _, v := range dfv {
		dynamic_field[v.Name] = v.ValueText
	}
	return dynamic_field
}

func TicketHistoryAdd() {

}

func TicketGet(ticketID int64) {

}
