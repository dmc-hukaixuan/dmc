package admin

import (
	"dmc/global"
	model "dmc/kernel/model/admin"
	"fmt"
)

type DynamicField struct{}

//var DynamicFieldA = new(DynamicField)

/*
	@param: fieldType string
	@description: get all dynamicfield detail
	@return: an array struct of model.dynmaicfield
*/
func DynamicFieldListGet() (dfenter []model.DynamicField, err error) {

	// var df []model.DynamicField
	selectSQL := `SELECT df.id AS id, df.internal_field AS internal_field , df.name AS name, df.label AS label,
					df.field_type AS field_type, df.object_type AS object_type,
					df.valid_id AS valid_id, df.create_time AS create_time,
					df.create_by AS create_by, df.change_time AS change_time, df.change_by AS change_by,
					u.full_name AS create_by_name, u1.full_name AS change_by_name
					FROM dynamic_field df
					LEFT JOIN users u ON u.id = df.create_by
					LEFT JOIN users u1 ON u1.id = df.change_by`
	err = global.GVA_DB.Raw(selectSQL).Scan(&dfenter).Error
	if err != nil {
		return
	}
	return dfenter, err
}

/*
	for ticket template or other page
*/
func DynamicFieldList(fieldType string) (dfenter []model.DynamicField, err error) {

	// var df []model.DynamicField
	selectSQL := `SELECT id, internal_field, name, label, field_type, object_type,
				config, valid_id, create_time, create_by, change_time, change_by FROM dynamic_field WHERE object_type = ? AND valid_id = 1`
	err = global.GVA_DB.Raw(selectSQL, fieldType).Scan(&dfenter).Error
	if err != nil {
		panic(err)
	}
	// fieldList := map[string]string{}
	// for _, v := range df {
	// 	fieldList["DynamicField_"+v.Name] = v.Label
	// }
	return dfenter, err
}

// add new Dynamic Field config
func DynamicFieldAdd(df model.DynamicField) (df_id int, err error) {
	err = global.GVA_DB.Table("dynamic_field").Create(&df).Error
	if err != nil {
		panic(err)
	}
	return df.ID, err
}

// get Dynamic Field attributes
func DynamicFieldGet(df_id int, name string) (df model.DynamicField) {
	if df_id > 0 {
		err := global.GVA_DB.Table("dynamic_field").Where("id = ?", df_id).First(&df).Error
		if err != nil {
			panic(err)
		}
	} else {
		err := global.GVA_DB.Table("dynamic_field").Where("name = ?", name).First(&df).Error
		if err != nil {
			panic(err)
		}
	}

	return df
}

// update Dynamic Field content into database
func DynamicFieldUpdate(df model.DynamicField) (district_id int, err error) {
	fmt.Println("df ", df.ConfigT, " Config ", df.Config)
	err = global.GVA_DB.Table("dynamic_field").Model(&df).Omit("create_by", "create_time").Updates(df).Error
	if err != nil {
		return
	}
	return df.ID, err
}

// update Dynamic Field content into database
func DynamicFieldNameList(fieldType string) map[string]model.DynamicField {
	var dfenter []model.DynamicField
	selectSQL := `SELECT id, internal_field, name, label, field_type, object_type,
				config, valid_id, create_time, create_by, change_time, change_by FROM dynamic_field WHERE object_type = ? AND valid_id = 1`
	err := global.GVA_DB.Raw(selectSQL, fieldType).Scan(&dfenter).Error
	if err != nil {
		//panic(err)
	}
	dfm := map[string]model.DynamicField{}
	// build a string for
	for _, v := range dfenter {
		dfm[v.Name] = v
	}
	return dfm
}

//delete a Dynamic field entry. You need to make sure that all values are
//deleted before calling this function, otherwise it will fail on DBMS which check
//referential integrity.
func DynamicFieldDelete() {

}
