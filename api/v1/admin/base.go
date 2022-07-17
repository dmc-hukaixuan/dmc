package admin

// BaseController wraps common methods for controllers to host API
// type BaseController struct {
//     // C *gin.Context
//     // User      string
//     //audit     models.Audit
//     //UserModel *models.User
// }

type Admin struct {
    ProcessManagementApi
    TicketTemplateApi
    RoleApi
    DynamicFieldApi
    ServiceApi
    SLAApi
    WorkingTimeApi
    TicketTypeAPI
    TicketPriorityAPI
    TicketStateAPI
    DepartmentApi
    TicketSourceAPI
}

// check request is a json string
// func (b *BaseController) CheckData() {

// }

// // GetStringFromQuery gets the param from query and returns it as string
// func (b *BaseController) GetStringFromQuery(key string) string {
//     return b.Ctx.Input.Query(key)
// }
