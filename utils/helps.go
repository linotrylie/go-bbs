package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/netutil"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 类型转换
func InterfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

func Ip2long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)

	return
}

func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

func MaskPhone(phone string) string {
	if len(phone) <= 10 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}

func DatetimeToUnix(datetime string) int64 {
	loc, _ := time.LoadLocation("Local")
	parseInLocation, err := time.ParseInLocation(time.DateTime, datetime, loc)
	if err != nil {
		return 0
	}
	return parseInLocation.Unix()
}

func TrimHtml(src string) string {
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
	return strings.TrimSpace(src)
}

func GetIpCity(ip string) (string, error) {
	if ip == "" {
		return "", errors.New("invalid target param")
	}
	request := &netutil.HttpRequest{
		RawURL: "https://whois.pconline.com.cn/ipJson.jsp?ip=" + ip,
		Method: "GET",
	}
	httpClient := netutil.NewHttpClient()
	resp, err := httpClient.SendRequest(request)
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("invalid target param")
	}
	if resp == nil {
		return "", errors.New("invalid target param")
	}
	type IpResponse struct {
		Ip          string `json:"ip,omitempty"`
		Pro         string `json:"pro,omitempty"`
		ProCode     string `json:"proCode,omitempty"`
		City        string `json:"city,omitempty"`
		CityCode    string `json:"cityCode,omitempty"`
		Region      string `json:"region,omitempty"`
		RegionCode  string `json:"regionCode,omitempty"`
		Addr        string `json:"addr,omitempty"`
		RegionNames string `json:"regionNames,omitempty"`
		Err         string `json:"err,omitempty"`
	}
	var ipresp IpResponse
	defer resp.Body.Close()
	readAll, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	toUtf8, err := GbkToUtf8(readAll)
	if err != nil {
		return "", err
	}
	str := strings.TrimSpace(string(toUtf8))
	strN1 := strings.Replace(str, "if(window.IPCallBack) {IPCallBack(", "", 1)
	strN2 := strings.Replace(strN1, ");}", "", 1)
	err = json.Unmarshal([]byte(strN2), &ipresp)
	if err != nil {
		return "", err
	}
	return ipresp.Addr, nil
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
