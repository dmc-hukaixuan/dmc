package ticket

// BaseController wraps common methods for controllers to host API
// type BaseController struct {
//     // C *gin.Context
//     // User      string
//     //audit     models.Audit
//     //UserModel *models.User
// }

type TicketAPI struct {
	TemplateAPI
}

// check request is a json string
// func (b *BaseController) CheckData() {

// }

// // GetStringFromQuery gets the param from query and returns it as string
// func (b *BaseController) GetStringFromQuery(key string) string {
//     return b.Ctx.Input.Query(key)
// }
