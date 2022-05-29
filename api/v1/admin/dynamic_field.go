package admin

import (
	"dmc/kernel/service/admin"
	"fmt"

	"github.com/gin-gonic/gin"
)

func base(c *gin.Context) {

}

// get dynmaic field list
func DynamicFieldList(fieldType string) {

	// get field from db
	df, _ := admin.DynamicFieldA.DynamicFieldList(fieldType)
	for _, v := range df {
		fmt.Println("v : ", v.Label)
	}
}

func DynamicFieldGet() {

}
