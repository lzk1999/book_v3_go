package utils

import (
	"book_v3_go/setting"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/axgle/mahonia"
	"math"
	"net/http"
	"strconv"
)

func PassSalt(pass string, salt string) string {
	return MD5(((MD5(pass)) + salt))
}

func Img(group int64, articleid int64) string {
	return fmt.Sprintf("%s/%d/%d/%ds.jpg", setting.Conf.WebConfig.ImgUrl, group, articleid, articleid)
}

func Txt(id int64, cid int64) (result string, err error) {
	url := fmt.Sprintf("%s/%d/%d/%d.txt", setting.Conf.WebConfig.TxtUrl, int(math.Ceil(float64(id/10000))), id, cid)
	//请求
	result, err = httpGet(url)
	return result, err
}

// GBK转utf8的方法
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func httpGet(url string) (result string, err error) {
	// 借助 http包的 Get()函数 获取网页数据
	resp, err := http.Get(url)
	if resp.StatusCode == 404 {
		return "", err
	}
	if err != nil {
		// 将错误传出
		return
	}
	// 读取结束，关闭resp.Body
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		// 读取Body内容
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		// 拼接每次buf中读到的数据，到result中，返回
		result += ConvertToString(string(buf[:n]), "GBK", "UTF-8")
	}
	return
}

func TableName(id string) string {
	tid, _ := strconv.Atoi(id)
	return fmt.Sprintf("shipsay_article_chapter_%d", int(math.Ceil(float64(tid/10000))+1))
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
