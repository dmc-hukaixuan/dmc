/*
* All type functions.
 */
package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/ticket"
	"fmt"
	"strconv"
)

func TypeAdd() {

}

func TypeGet() {

}

func TypeUpdate() {

}

/*
get type list

    my %List = $TypeObject->TypeList();

or

    my %List = $TypeObject->TypeList(
        Valid => 0,
    );
*/
func TypeList(validID int) map[string]string {
	var tp []model.TicketType
	selectSQL := `SELECT id, name FROM ticket_type`
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
