package ticket

import (
    "dmc/kernel/model/common/response"
    "dmc/kernel/service/ticket"

    "github.com/gin-gonic/gin"
)

type TicketSearch struct{}

// ��������
func (*TicketSearch) SearchCondition(c *gin.Context) {
    user_id, _ := c.Get("userID")
    // get user profile
    userProfileList := ticket.SearchProfileList(user_id.(int))

    // get ticket search field data

    response.SuccessWithDetailed(gin.H{
        "userProfileList": userProfileList,
        "fieldData":       "",
        "fieldOrder":      "",
        "defaultField":    "",
    }, "��ȡ�ɹ�", c)
}

func SearchPorfileGet() {

}
