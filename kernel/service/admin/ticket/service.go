/*
All service functions.
*/
package ticket

import (
	"dmc/global"
	"fmt"
	"strconv"

	model "dmc/kernel/model/ticket"
)

func ServiceList(validID int) map[string]string {
	var sla []model.Service
	selectSQL := `SELECT id, name FROM service`
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

func ServiceGet() {

}

func ServiceAdd() {

}

func ServiceUpdate() {

}

func ServiceLinkSLA() {

}
