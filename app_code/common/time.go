package common

import (
	"fmt"
	"time"
)

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

//获取到当前日期和时间,按照北京时间计算
func GetBeiJingTime() string {
	t, _ := TimeIn(time.Now(), "Asia/Shanghai")
	return fmt.Sprintf("%s", t.Format("2006-01-02 15:04:05"))
}

// 获取当前时间戳
func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

func GetTime_Now_Year() int {
	return time.Now().Year()
}

func GetTime_Now_Month() int {
	return int(time.Now().Month())
}

func GetTime_Now_Day() int {
	return time.Now().Day()
}

// 获取指定时间的时间戳
// 例 GetTimestamp("2018-06-07 12:00:00")
func GetTimestamp(t string) (err error, timestamp int64) {
	var formatTime time.Time
	if formatTime, err = time.Parse(t, "2018-06-07 12:00:00"); err != nil {
		timestamp = 0
		return
	}

	timestamp = formatTime.Unix()
	return
}

// 获取当前一段时间后的时间戳
// 例 GetAfterNowTimestamp(0,)
func GetAfterNowTimestamp(hour, minute, second time.Duration) int64 {
	return time.Now().Add(hour*time.Hour + minute*time.Minute + second*time.Second).Unix()
}
