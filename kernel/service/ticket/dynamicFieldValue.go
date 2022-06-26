package ticket

import (
	"dmc/global"
	model "dmc/kernel/model/dynamicField"
)

// Delete all entries of a dynamic field values for object ID.
func ObjectValuesDelete(objectType string, objectID int) {
	deleteSQL := ` DELETE FROM dynamic_field_value
    WHERE
        field_id IN (
            SELECT id FROM dynamic_field
            WHERE object_type = ?
        )
        AND object_id = ?`
	global.GVA_DB.Raw(deleteSQL, objectType, objectID).Unscoped()
}

/*
get all distinct values from a field stored on the database

    my $HistoricalValues = $DynamicFieldValueObject->HistoricalValueGet(
        FieldID   => $FieldID,                  # ID of the dynamic field
        ValueType => 'Text',                    # or 'DateTime' or 'Integer'. Default 'Text'
    );

    Returns:

    $HistoricalValues{
        ValueA => 'ValueA',
        ValueB => 'ValueB',
        ValueC => 'ValueC'
    };
*/
func HistoricalValueGet() {

}

/*
delete all entries of a dynamic field .

    my $Success = $DynamicFieldValueObject->AllValuesDelete(
        FieldID            => $FieldID,                 # ID of the dynamic field
        UserID  => 123,
    );

    Returns 1.
*/
func AllValuesDelete(fieldID int) {
	global.GVA_DB.Raw("DELETE FROM dynamic_field_value WHERE field_id = ?", fieldID).Unscoped()
	// Cleanup entire cache!

}

/*
delete a Dynamic field value entry. All associated rows will be deleted.

    my $Success = $DynamicFieldValueObject->ValueDelete(
        FieldID            => $FieldID,                 # ID of the dynamic field
        ObjectID           => $ObjectID,                # ID of the current object that the field
                                                        #   is linked to, e. g. TicketID
        UserID  => 123,
    );

    Returns 1.
*/
func ValueDelete(fieldID, objectID int) {
	// delete dynamic field value
	global.GVA_DB.Raw("DELETE FROM dynamic_field_value WHERE field_id = ? AND object_id = ?", fieldID, objectID).Unscoped()
	// should be to do delete cache
}

/*
get a dynamic field value. For each table row there will be one entry in the
result list.

    my $Value = $DynamicFieldValueObject->ValueGet(
        FieldID            => $FieldID,                 # ID of the dynamic field
        ObjectID           => $ObjectID,                # ID of the current object that the field
                                                        #   is linked to, e. g. TicketID
    );

    Returns [
        {
            ID                 => 437,
            ValueText          => 'some text',
            ValueDateTime      => '1977-12-12 12:00:00',
            ValueInt           => 123,
        },
    ];
*/
func ValueGet(fieldID, objectID int) (values []model.DynamicFieldValue) {
	// get cache object

	// Special caching strategy: cache all fields of an object in one cache file.
	// 	This avoids too many cache files on systems with many fields for many objects.
	selectSQL := `SELECT id, value_text, value_date, value_int, field_id
                    FROM dynamic_field_value
                    WHERE object_id = ?
                    ORDER BY id`

	global.GVA_DB.Raw(selectSQL, objectID).Scan(&values)
	// We'll populate cache with all object's dynamic fields to reduce
	// number of db accesses (only one db query for all dynamic fields till
	// cache expiration); return only specified one dynamic field
	return values
}

func ValueSet(fieldID int, objectID int, values []model.DynamicFieldValue) {
	//delete existing value
	ValueDelete(fieldID, objectID)
	// create a new value entry
	err := global.GVA_DB.Table("dynamic_field_value").Create(&values).Error
	if err != nil {
		panic(err)
	}
	// delete cache

}
