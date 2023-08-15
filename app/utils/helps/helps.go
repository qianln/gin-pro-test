package helps

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gin-pro/app/core/system"
	"gin-pro/app/global/consts"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func MD5(params string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(params))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func Base64Md5(params string) string {
	return MD5(base64.StdEncoding.EncodeToString([]byte(params)))
}

// GetFilesMimeByFileName 通过文件名获取文件mime信息
func GetFilesMimeByFileName(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		system.ZapLog.Error(consts.ErrorsFilesUploadOpenFail + err.Error())
	}
	defer f.Close()

	// 只需要前 32 个字节就可以了
	buffer := make([]byte, 32)
	if _, err := f.Read(buffer); err != nil {
		system.ZapLog.Error(consts.ErrorsFilesUploadReadFail + err.Error())
		return ""
	}

	return http.DetectContentType(buffer)
}

// GetFilesMimeByFp 通过文件指针获取文件mime信息
func GetFilesMimeByFp(fp multipart.File) string {

	buffer := make([]byte, 32)
	if _, err := fp.Read(buffer); err != nil {
		system.ZapLog.Error(consts.ErrorsFilesUploadReadFail + err.Error())
		return ""
	}

	return http.DetectContentType(buffer)
}

type Err struct {
	Value string
}

// 实现error接口
func (e Err) Error() string {
	return e.Value
}

func GetIntFromChar(value string) (int64, error) {
	val := value[2:]
	n, err := strconv.ParseInt(val, 16, 32)
	return n, err
}

func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func LocalLog(value any) {
	defer func() {
		if recover() != nil {
			fmt.Println("打印日志异常")
		}
	}()
	switch value.(type) {
	case int:
		system.ZapLog.Info(fmt.Sprintf("%d", value))
	case int64:
		system.ZapLog.Info(fmt.Sprintf("%d", value))
	case string:
		system.ZapLog.Info(fmt.Sprintf("%v", value))
	case []int:
		marshal, err := json.Marshal(value)
		if err != nil {
			fmt.Println("打印日志异常")
			return
		}
		system.ZapLog.Info(fmt.Sprintf("%v", string(marshal)))
	case []string:
		marshal, err := json.Marshal(value)
		if err != nil {
			fmt.Println("打印日志异常")
			return
		}
		system.ZapLog.Info(fmt.Sprintf("%v", string(marshal)))
	default:
		marshal, err := json.Marshal(value)
		if err != nil {
			fmt.Println("打印日志异常")
			return
		}
		system.ZapLog.Info(fmt.Sprintf("%v", string(marshal)))
	}
}
