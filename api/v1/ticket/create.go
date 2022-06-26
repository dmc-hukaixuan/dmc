package ticket

import(
	"dmc/kernel/util/uploadCache"
	model "dmc/kernel/model/ticket"
	service "dmc/kernel/service/ticket"
	dynamicFieldObject "dmc/kernel/service/template"
)

func TicketTemplateList() {

}

func TicketTemplateGet() {
	// get formid
	FormID := uploadCache.WebUploadCache().FormIDCreate()

	
	ticketID := $TicketObject->TicketCreate(
		Title        => $GetParam{Subject},
		QueueID      => $NewQueueID,
		Subject      => $GetParam{Subject},
		Lock         => 'unlock',
		TypeID       => $GetParam{TypeID},
		ServiceID    => $GetParam{ServiceID},
		SLAID        => $GetParam{SLAID},
		StateID      => $GetParam{NextStateID},
		PriorityID   => $GetParam{PriorityID},
		OwnerID      => 1,
		CustomerNo   => $CustomerID,
		CustomerUser => $SelectedCustomerUser,
		UserID       => $Self->{UserID},
	);

	
	DYNAMICFIELD:
	for my $DynamicFieldConfig ( @{ $Self->{DynamicField} } ) {
		next DYNAMICFIELD if !IsHashRefWithData($DynamicFieldConfig);
		next DYNAMICFIELD if $DynamicFieldConfig->{ObjectType} ne 'Ticket';

		# set the value
		my $Success = $DynamicFieldBackendObject->ValueSet(
			DynamicFieldConfig => $DynamicFieldConfig,
			ObjectID           => $TicketID,
			Value              => $DynamicFieldValues{ $DynamicFieldConfig->{Name} },
			UserID             => $Self->{UserID},
		);
	}
}

func TicketCreate(c *gin.Context) {
	var ticketBaseData model.TicketBaseData
	_ = c.ShouldBindJSON(&ticketBaseData)
	// create new ticket, do db insert
	ticket_id, _ := service.TicketCreate(ticketBaseData)
	
	ticketData := make(map[string]interface{}) //注意该结构接受k的内容
    c.BindJSON(&ticketData)

	// get all dynamicField id
	DynamicField := dynamicField.DynamicFieldNameList("Ticket")
	reg := regexp.MustCompile(`^dynamicField_`)
	// set ticket dynamic fields
	// cycle through the activated Dynamic Fields for this screen
	for k, v := range ticketData {
		if reg.MatchString(k) {
			//  set the value
			dynamicFieldObject.DynamicField(DynamicField[k].FieldType).ValueSet(fieldID, objectID, value)
		}
	}

	// article create
	// get pre loaded attachment
	
}