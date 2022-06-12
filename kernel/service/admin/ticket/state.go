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

/*
	type list get
*/
func StateListGet(validID int) (tt []model.TicketState) {
	selectSQL := `SELECT tt.id, tt.name, tt.valid_id, tst.name AS state_type, tt.comments AS comments,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					tt.create_time AS create_time, tt.change_time AS change_time
					FROM ticket_state tt 
					LEFT JOIN users u ON u.id = tt.create_by 
					LEFT JOIN ticket_state_type  tst ON tst.id = tt.id
					LEFT JOIN users u1 ON u1.id = tt.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&tt).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return tt
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
