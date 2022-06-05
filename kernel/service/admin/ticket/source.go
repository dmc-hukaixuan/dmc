package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

/*
get type list

    my %List = $TypeObject->TypeList();

or

    my %List = $TypeObject->TypeList(
        Valid => 0,
    );
*/
func SourceList(validID int) map[string]string {
	var tp []model.TicketType
	selectSQL := `SELECT id, name FROM ticket_source`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	typeList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&tp).Error
	if err != nil {
		return typeList
	}
	// do loop, build a json string
	for _, v := range tp {
		typeList[strconv.Itoa(v.ID)] = v.Name
	}
	return typeList
}
