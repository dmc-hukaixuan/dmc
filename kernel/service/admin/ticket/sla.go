package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

func SLAGet() {

}

func SLAList(validID int) map[string]string {
	var sla []model.SLA
	selectSQL := `SELECT id, name FROM sla`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	slaList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&sla).Error
	if err != nil {
		return slaList
	}
	// do loop, build a json string
	for _, v := range sla {
		slaList[strconv.Itoa(v.ID)] = v.Name
	}
	return slaList
}
