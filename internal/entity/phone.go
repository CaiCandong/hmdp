package entity

import (
	"regexp"
)

type Phone string

func (p Phone) VerifyMobileFormat() bool {
	// 校验手机号码格式
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(string(p))
}

//func (p Phone) GenValidateCode(width int) string {
//	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	r := len(numeric)
//	rand.Seed(time.Now().UnixNano())
//
//	var sb strings.Builder
//	for i := 0; i < width; i++ {
//		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
//	}
//	return sb.String()
//}
