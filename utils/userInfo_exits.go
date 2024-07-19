package utils

import (
	"net"
	"regexp"
	"strings"
)

// 邮箱加密
func DesensitizationEmail(email string) string {
	emailList := strings.Split(email, "@")
	return emailList[0][:4] + "***@" + emailList[1]
}

// 手机号码加密
func DesensitizationPhone(phone string) string {
	if len(phone) != 11 {
		return ""
	}
	return phone[:3] + "****" + phone[8:]
}

// 邮箱检验
func IsValidEmail(email string) bool {
	// 正则表达式检查邮箱格式
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	if !re.MatchString(email) {
		return false
	}

	// 检查域名是否有效
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false
	}

	return true
}
