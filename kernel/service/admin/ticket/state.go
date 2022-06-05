/*
* All ticket state functions.
 */
package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

/*
 add a ticket state
*/
func StateAdd(ts model.TicketState) (stateID int, err error) {
	err = global.GVA_DB.Table("ticket_state").Create(&ts).Error
	if err != nil {
		return
	}
	return ts.ID, err
}

/*
	ticket state detial info
*/
func StateGet(stateID int) (ts model.TicketState) {
	err := global.GVA_DB.Table("ticket_state").Where("id = ?", stateID).First(&ts).Error
	if err != nil {
		return
	}
	return ts
}

/*
	update ticket state
*/
func StateUpdate(ts model.TicketState) (stateID int, err error) {
	err = global.GVA_DB.Table("ticket_state").Where("id = ?", ts.ID).Model(&ts).Omit("create_by", "create_time").Updates(ts).Error
	if err != nil {
		return
	}
	return ts.ID, err
}

/*
* get state list as a hash of ID, Name pairs

useage

returns

	 %List = (
        1 => "new",
        2 => "closed successful",
        3 => "closed unsuccessful",
        4 => "open",
        5 => "removed",
        6 => "pending reminder",
        7 => "pending auto close+",
        8 => "pending auto close-",
        9 => "merged",
    );

*/
func StateList(validID int) map[string]string {
	var tp []model.TicketState
	selectSQL := `SELECT id, name FROM ticket_state`
	fmt.Print("validID", validID)
	if validID > 0 {
		selectSQL = fmt.Sprint(selectSQL, " WHERE valid_id = ", validID)
	}
	stateList := make(map[string]string)
	err := global.GVA_DB.Raw(selectSQL).Scan(&tp).Error
	if err != nil {
		return stateList
	}
	// do loop, build a json string
	for _, v := range tp {
		stateList[strconv.Itoa(v.ID)] = v.Name
	}
	return stateList
}

func StateGetStatesByType() {

}

/*

get state type list as a hash of ID, Name pairs

    my %ListType = $StateObject->StateTypeList(
        UserID => 123,
    );

returns

    my %ListType = (
        1 => "new",
        2 => "open",
        3 => "closed",
        4 => "pending reminder",
        5 => "pending auto",
        6 => "removed",
        7 => "merged",
    );

*/
func StateTypeList() map[int]string {
	var tp []model.TicketStateType
	err := global.GVA_DB.Table("ticket_state_type").Scan(&tp).Error
	stateTypeList := map[int]string{}
	if err != nil {
		return stateTypeList
	}
	for _, v := range tp {
		stateTypeList[v.ID] = v.Name
	}
	return stateTypeList
}
