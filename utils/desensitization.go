package utils

import (
	"strings"
)

func DesensitizationPhone(phone string) string {
	if len(phone) != 11 {
		return ""
	}
	return phone[:3] + "****" + phone[8:]
}

func DesensitizationEmail(email string) string {
	emailList := strings.Split(email, "@")
	return emailList[0][:4] + "***@" + emailList[1]
}
