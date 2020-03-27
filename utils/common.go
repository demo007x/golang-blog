package utils

import (
	"blog/config"
	"blog/modules"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// md5 func
func Md5(str string) string {
	data := []byte(str)
	rest := fmt.Sprintf("%x", md5.Sum(data))

	return rest
}

// md5 fun v2
func md5V2(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 加密用户密码
func CryptUserPassword(password string, salt string) string {
	return Md5(password + salt)
}

// 获取项目配置数据
func Config(key string) interface{} {
	return nil
}

// 获取4位密码的盐值
func Salt() string {
	rand.Seed(time.Now().UnixNano()) // 伪随机种子
	baseStr := "abcdefghigklmnopqistuvwxyzABCDEFGHIGKLMNOPQISTUVWXYZ0123456789"
	saltLen := 4
	salt := make([]byte, saltLen)
	for n := 0; n < saltLen; n++ {
		salt[n] = baseStr[rand.Int31n(int32(len(baseStr)))]
	}

	return string(salt)
}

// 验证用户的密码是否正确
func VerifyUserPassword(user *modules.User, oldPsd string) bool {
	password := CryptUserPassword(oldPsd, user.Salt)
	if password == user.Password {
		return true
	}

	return false
}

// Html 非转义输出html
func Html(str string) interface{} {
	return template.HTML(str)
}

// 将html转换string
func Html2Str(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	src = strings.TrimSpace(src)
	src = strings.Trim(src, "#")
	return strings.Replace(src, "\n", "", len([]rune(src)))
}

// TagString2Map 将文章tags转换为[]string
func TagString2Map(tagString string) []string {
	var tagSlice []string
	tagSlice = strings.Split(tagString, ",")

	for i, v := range tagSlice {
		tagSlice[i] = strings.Trim(v, " ")
	}

	return tagSlice
}

// SetLinkTitle 设置a链接的title
func SetLinkTitle(title string) string {
	section,err := config.Config.GetSection("env")
	if err != nil {
		return ""
	}

	appName,_ := section.GetKey("AppName")
	//appUrl,_ := section.GetKey("AppUrl")

	return fmt.Sprintf("%s - %s", title, appName)
}

// AppUrl 获取站点配置的url
func AppUrl(path string) string {
	section,err := config.Config.GetSection("env")
	if err != nil {
		return fmt.Sprintf("%s", path)
	}
	appUrl,err := section.GetKey("AppUrl")
	if err != nil {
		return fmt.Sprintf("%s", path)
	}
	return fmt.Sprintf("%s%s", appUrl, path)
}

// SocialHtml 返回社交html
func SocialHtml() string  {
	var socialHtml string

	config := config.Config
	socialSection,err := config.GetSection("social")
	if err != nil {
		return socialHtml
	}

	// 定义映射关系
	socialMap := make(map[string]string)
	socialMap["github"] = `<i class="fab fa-github" aria-hidden="true"></i>`
	socialMap["twitter"] = `<i class="fab fa-twitter" aria-hidden="true"></i>`
	socialMap["linkedin"] = `<i class="fab fa-linkedin" aria-hidden="true"></i>`
	socialMap["stack"] = `<i class="fab fa-stack-overflow" aria-hidden="true"></i>`
	socialMap["codepen"] = `<i class="fab fa-codepen" aria-hidden="true"></i>`

	// <li class="list-inline-item"><a href="#"><i class="fab fa-twitter fa-fw"></i></a></li>

	for _, key := range socialSection.KeyStrings() {
		val, _ := socialSection.GetKey(key)
		if keyHtml,ok := socialMap[key]; ok {
			socialHtml = fmt.Sprintf(`%s<li class="list-inline-item"><a href="%s">%s</a></li>`, socialHtml, val, keyHtml)
		} else {
			socialHtml = fmt.Sprintf(`%s<li class="list-inline-item"><a href="%s">%s</a></li>`, socialHtml, val, key)
		}
	}

	return socialHtml
}
