package util

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

func GenerateCode() string {
	code := GenerateUid()
	parts := strings.Split(code, "-")
	if len(parts) < 4 {
		return code
	}
	result := parts[2] + "-" + parts[3]

	var letterResult strings.Builder
	for _, char := range result {
		if char != '-' {
			letterResult.WriteRune(NumToLetter(char))
		} else {
			letterResult.WriteRune(char) // 保留 -
		}
	}
	return strings.ToUpper(letterResult.String())
}

func NumToLetter(num rune) rune {
	return 'A' + (num - '0') // 将数字转换为字母 A-J
}

// GenerateUid generates a UUIDv7 and retries up to 10 times in case of error.
func GenerateUid(length ...int64) string {
	var uid uuid.UUID
	var err error

	for i := 0; i < 10; i++ {
		if uid, err = uuid.NewV7(); err == nil {
			return uid.String()
		}
	}

	return fmt.Sprintf("%d", time.Now().UnixNano())
}

//f47ac10b-58cc-0372-8567-0e02b2c3d479

func GenerateOrderSN(memberId int64) string {
	formattedTime := time.Now().Format("060102150405")

	return fmt.Sprintf("%s%07d", formattedTime, memberId)
}

func GenerateRefundSN(orderID, orderItemID int64) string {

	return fmt.Sprintf("%s%07d%07d", "Refund", orderID, orderItemID)
}
