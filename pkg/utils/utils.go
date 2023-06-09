package utils

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func UUID() string {
	u := uuid.New()
	return u.String()
}

// ToJSON 辅助函数 - 将结构体序列化为 JSON 字符串
func ToJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// FromJSON 辅助函数 - 将 Redis 响应解码为结构体
func FromJSON(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}

// SnowFlakeID 生成雪花ID
func SnowFlakeID() uint {
	// 41位时间戳
	t := time.Now().UnixNano() / 1e6
	// 10位机器ID
	mid := int64(1)
	// 12位序列号
	seq := int64(1)
	return uint((t-1600000000000)<<22 | mid<<12 | seq)
}

// GenOrderNo 生成订单号
func GenOrderNo() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 4; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	// 13位时间戳 精确到毫秒
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	return timestamp + sb.String()
}

func IsInTime(startTime, endTime time.Time) bool {
	now := time.Now()
	return now.After(startTime) && now.Before(endTime)
}
