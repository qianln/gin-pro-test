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
	"strings"
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

/**
 * 驼峰转蛇形 snake string
 * @description XxYy to xx_yy , XxYY to xx_y_y
 * @date 2020/7/30
 * @param s 需要转换的字符串
 * @return string
 **/
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 * @date 2020/7/30
 * @param s要转换的字符串
 * @return string
 **/
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// CreateFile 创建文件
func CreateFile(basePath, fileName string, content string) error {
	if err := os.MkdirAll(basePath, 0755); err != nil {
		system.ZapLog.Error("创建文件夹失败")
		return err
	}
	_, err := os.Stat(basePath + "/" + fileName + ".go")
	if err == nil {
		system.ZapLog.Warn("文件存在")
		return fmt.Errorf("文件存在")
	}
	file, err := os.Create(basePath + "/" + fileName + ".go")
	if err != nil {
		system.ZapLog.Error("创建文件失败")
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	if err != nil {
		system.ZapLog.Error("写入文件失败")
		return err
	}
	system.ZapLog.Info("创建一个控制器")
	return nil
}
