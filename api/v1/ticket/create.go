package ticket

import (
	model "dmc/kernel/model/ticket"
	dynamicField "dmc/kernel/service/admin"
	userService "dmc/kernel/service/admin/user"
	dynamicFieldObject "dmc/kernel/service/template"
	service "dmc/kernel/service/ticket"
	article "dmc/kernel/service/ticket/article"
	"dmc/kernel/util/uploadCache"
	"regexp"

	"github.com/gin-gonic/gin"
)

func TicketTemplateList() {

}

func TicketTemplateGet() {
	// get formid
	// FormID := uploadCache.WebUploadCache().FormIDCreate()

	// ticketID := $TicketObject->TicketCreate(
	// 	Title        => $GetParam{Subject},
	// 	QueueID      => $NewQueueID,
	// 	Subject      => $GetParam{Subject},
	// 	Lock         => 'unlock',
	// 	TypeID       => $GetParam{TypeID},
	// 	ServiceID    => $GetParam{ServiceID},
	// 	SLAID        => $GetParam{SLAID},
	// 	StateID      => $GetParam{NextStateID},
	// 	PriorityID   => $GetParam{PriorityID},
	// 	OwnerID      => 1,
	// 	CustomerNo   => $CustomerID,
	// 	CustomerUser => $SelectedCustomerUser,
	// 	UserID       => $Self->{UserID},
	// );

	// DYNAMICFIELD:
	// for my $DynamicFieldConfig ( @{ $Self->{DynamicField} } ) {
	// 	next DYNAMICFIELD if !IsHashRefWithData($DynamicFieldConfig);
	// 	next DYNAMICFIELD if $DynamicFieldConfig->{ObjectType} ne 'Ticket';

	// 	# set the value
	// 	my $Success = $DynamicFieldBackendObject->ValueSet(
	// 		DynamicFieldConfig => $DynamicFieldConfig,
	// 		ObjectID           => $TicketID,
	// 		Value              => $DynamicFieldValues{ $DynamicFieldConfig->{Name} },
	// 		UserID             => $Self->{UserID},
	// 	);
	// }
}

func TicketCreate(c *gin.Context) {
	var ticketBaseData model.TicketBaseData
	_ = c.ShouldBindJSON(&ticketBaseData)

	user_id, _ := c.Get("userID")
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
			dynamicFieldObject.DynamicField(DynamicField[k].FieldType).ValueSet(DynamicField[k].ID, "Ticket", ticket_id, v)
		}
	}

	form := userService.UserGet(ticketData["customer"].(int))
	to := userService.RoleGet(ticketData["queue"].(int))

	SenderType := "agent"

	// article create
	// get pre loaded attachment
	articleID := article.ArticleCreate(
		model.ArticleDataMimeCreate{
			TicketID:             ticket_id,
			SenderType:           SenderType,
			IsVisibleForCustomer: ticketData["isVisibleForCustomer"].(int),
			From:                 form.FullName + " <" + form.Email + " >",
			To:                   to.Name,
			Subject:              ticketData["subject"].(string),
			Body:                 ticketData["body"].(string),
			MimeType:             "text/plain",
			Charset:              "charset=utf-8",
			UserID:               user_id.(int),
		},
	)
	// write attachments and form id
	article.ArticleWriteAttachment(ticketData["formID"].(string), articleID, user_id.(int), ticket_id)

	// remove pre submited attachments
	uploadCache.WebUploadCache.FormIDRemoveFile(ticketData["formID"].(string))
	
	my $Success = $ProcessObject->ProcessTicketProcessSet(
		ProcessEntityID => $Param{ProcessEntityID},
		TicketID        => $TicketID,
		UserID          => $Self->{UserID},
	);

	$Success = $ProcessObject->ProcessTicketActivitySet(
		ProcessEntityID  => $Param{ProcessEntityID},
		ActivityEntityID => $ProcessStartpoint->{Activity},
		TicketID         => $TicketID,
		UserID           => $Self->{UserID},
	);
	// trgger event 

	// link tickets
	// SplitLinkType， LinkType， Direction
	if ticketData['splitTicketID'] {

	}
	// check link ticket permission

}
