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
 */
func StateAdd() {

}

func StateGet() {

}

func StateUpdate() {

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
