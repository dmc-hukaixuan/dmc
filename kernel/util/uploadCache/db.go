package uploadCache

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type DB struct{}

func (*DB) FormIDCreate() string {
	rand.Seed(time.Now().UnixNano())
	// return requested form id
	return time.Now().GoString() + "." + string(rand.Intn(12341241))
}

func (*DB) FormIDRemove(formid string) {

}

func (*DB) FormIDAddFile(context *gin.Context) {

}

func (*DB) FormIDRemoveFile(formid string) {

}

func (*DB) FormIDGetAllFilesData(formid string) {

}

func (*DB) FormIDGetAllFilesMeta(formid string) {

}

func (*DB) FormIDCleanUp(formid string) {

}
