package admin

import (
	"dmc/global"
	model "dmc/kernel/model/admin"
)

type DynamicField struct{}

var DynamicFieldA = new(DynamicField)

func (d *DynamicField) DynamicFieldList(fieldType string) (dfenter []model.DynamicField, err error) {
	// var df []model.DynamicField
	selectSQL := `SELECT id, internal_field, name, label, field_type, object_type,
				config, valid_id, create_time, create_by, change_time, change_by FROM dynamic_field WHERE object_type = ?`
	err = global.GVA_DB.Raw(selectSQL, fieldType).Scan(&dfenter).Error
	if err != nil {
		return
	}
	return dfenter, err
}

func (d *DynamicField) DynamicFieldGet() {

}

func (d *DynamicField) DynamicFieldAdd() {
	
}

func (d *DynamicField) DynamicFieldUpdate() {
	
}

func (d *DynamicField) DynamicFieldDelete() {
	
}
