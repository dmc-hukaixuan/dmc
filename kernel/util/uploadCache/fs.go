package uploadCache

import (
	"dmc/global/log"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FS struct{}

func (*FS) FormIDCreate() string {
	end := strconv.Itoa(rand.Intn(12341241))
	rand.Seed(time.Now().UnixNano())
	// return requested form id
	return strconv.FormatInt(time.Now().UnixMicro(), 10) + "." + end
}

func (*FS) FormIDRemove(formid string) {

}

func (*FS) FormIDAddFile(context *gin.Context) {
	formid := context.PostForm("formid")
	// check formid is valid
	if !FormIDValidate(formid) {
		log.Logger.Error("FormID already exists, Please refresh page.")
	}

	// create cache subdirectory if not exist
	// Directory := global.CONFIG.Local.Path + '/' + formid
	Directory := "/var/temp/web_upload_cache/" + formid
	file, err := context.FormFile("file")
	if err != nil {
		log.Logger.Error("ERROR: upload file failed . %s", log.Any("serverError", err))
	}
	dst := fmt.Sprintf(Directory + file.Filename)
	// 保存文件至指定路径
	err = context.SaveUploadedFile(file, dst)

	if err != nil {
		log.Logger.Error("ERROR: save file failed", log.Any("serverError", err.Error()))
	}
}

type WebCacheUpload struct {
	FormID   string `json:"formID"`
	Filename string `json:"filename"`
}

/**
formid
key : fileName
*/
func (*FS) FormIDRemoveFile(formid string) {
	//p := global.CONFIG.Local.Path + "/" + formid + "/"
	p := "/var/temp/web_upload_cache/" + formid + "/"
	if strings.Contains(p, "/var/temp/web_upload_cache/") {
		if err := os.Remove(p); err != nil {
			log.Logger.Error("本地文件删除失败", log.Any("serverError", err.Error()))
		}
	}
}

func (*FS) FormIDGetAllFilesData(formid string) {

}

func (*FS) FormIDGetAllFilesMeta(formid string) {

}

/*
	Removed no longer needed temporary files.

	Each file older than 1 day will be removed.
*/
func (*FS) FormIDCleanUp(formid string) {

	retentionTime := time.Now().Unix() - 86400
	// read file
	files, err := ioutil.ReadDir("TempDir") //读取目录下文件
	if err != nil {
		return
	}
	for _, subdir := range files {
		subdirTime := strings.Split(subdir.Name(), ".")
		n, err := strconv.ParseInt(subdirTime[0], 10, 64)
		if err != nil {
			log.Logger.Error("Get seesion fail", log.Any("serverError", err))
		}
		if retentionTime > n {
			os.RemoveAll("TempDir" + subdir.Name())
		}
	}
}

func FormIDValidate(formid string) bool {
	match, _ := regexp.MatchString("^d+.d+.d+$", formid)
	if !match {
		return false
	}
	return true
}
