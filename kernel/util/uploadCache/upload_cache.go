package uploadCache

import "github.com/gin-gonic/gin"

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type UploadCache interface {
	FormIDCreate() string
	FormIDRemove(formid string)
	FormIDAddFile(context *gin.Context)
	FormIDRemoveFile(formid string)
	FormIDGetAllFilesData(formid string)
	FormIDGetAllFilesMeta(formid string)
	FormIDCleanUp(formid string)
}

// 上传附件接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func WebUploadCache() UploadCache {
	moduleType := "fs"
	switch moduleType {
	case "fs":
		return &FS{}
	case "db":
		return &DB{}
	default:
		return &FS{}
	}
}
