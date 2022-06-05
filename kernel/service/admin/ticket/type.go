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

/*
	add a new ticket type
*/
func TypeAdd(tt model.TicketType) (type_id int, err error) {
	err = global.GVA_DB.Table("ticket_type").Create(&tt).Error
	if err != nil {
		return
	}
	return tt.ID, err
}

/*
	get types attributes
*/
func TypeGet(typeID int) (ts model.TicketType) {
	err := global.GVA_DB.Table("ticket_type").Where("id = ?", typeID).First(&ts).Error
	if err != nil {
		return
	}
	return ts
}

func TypeUpdate(tt model.TicketType) (typeID int, err error) {
	err = global.GVA_DB.Table("ticket_type").Where("id = ?", tt.ID).Model(&tt).Omit("create_by", "create_time").Updates(tt).Error
	if err != nil {
		return
	}
	return tt.ID, err
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

/*
	type list get
*/
func TypeListGet(validID int) (tt []model.TicketType) {
	selectSQL := `SELECT tt.id, tt.name, tt.valid_id, tt.tnstart AS tnstart,
					u.full_name AS create_by_name, u1.full_name AS change_by_name, 
					tt.create_time AS create_time, tt.change_time AS change_time
					FROM ticket_source tt 
					LEFT JOIN users u ON u.id = tt.create_by 
					LEFT JOIN users u1 ON u1.id = tt.change_by`
	err := global.GVA_DB.Raw(selectSQL).Scan(&tt).Error
	if err != nil {
		fmt.Println("err", err)
	}
	return tt
}
