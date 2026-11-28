package time_token

import (
	"printer/config"
	"printer/core/hash"

	"time"
)

const timeTokenSalt = "285a2da795fff7f73e7379e49d6289c8"
const tokenTimeFormat = "20060102150405" // 2006-01-02 15:04:05

func CreateToken() string {
	var timeStr = time.Now().Format(tokenTimeFormat)
	return timeStr + encryptTime(timeStr)
}

func CheckToken(token string) (bool, error) {
	var timeLength = len(tokenTimeFormat)
	if len(token) <= timeLength {
		return false, nil
	}

	// 检查有效性
	var timeStr = token[0:timeLength]
	var checkStr = token[timeLength:]

	if checkStr != encryptTime(timeStr) {
		return false, nil
	}

	// 检查时效性
	var timeInToken, err = time.ParseInLocation(tokenTimeFormat, timeStr, config.TimeLoc())
	if err != nil {
		// err = fmt.Errorf("解析时间发生错误，%w", err)
		// logrus.Error(err)
		// return false, err

		// 其实这个不算是错误，完全可以只算是 token 不合法
		return false, nil
	}

	// 暂定 5 分钟前后有效
	var now = time.Now()
	var validTimeBegin = now.Add(-5 * time.Minute)
	var validTimeEnd = now.Add(5 * time.Minute)

	if timeInToken.Before(validTimeBegin) {
		return false, nil
	}

	if timeInToken.After(validTimeEnd) {
		return false, nil
	}

	return true, nil
}

func encryptTime(timeStr string) string {
	return hash.SHA256(timeStr + timeTokenSalt)
}
