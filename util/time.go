package util

import (
	"fmt"
	"math"
	"time"
)

var Location *time.Location

func Init(locationName string) {
	l, err := time.LoadLocation(locationName)
	if err != nil {
		panic(err)
	}
	Location = l
}

func Now() time.Time {
	return time.Now().In(Location)
}

func ParseTime(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func DateTime(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, Location)
}

func GetTime(t *time.Time) string {
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
}

func GetWeekDay(t *time.Time) int {
	return int(t.Weekday())
}

func GetWeekNumber(t *time.Time) int {
	// 年間の週番号 - 月初の週番号 + 1
	return americanWeekNumber(t.Year(), int(t.Month()), t.Day()) - americanWeekNumber(t.Year(), int(t.Month()), 1) + 1
}

func americanWeekNumber(year, month, day int) int {
	// 指定した年の日付型作成
	newYearsDay := time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.Local)
	// 指定した月の日付型作成
	specifiedDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	// 曜日取得(1/1の曜日)
	newWeekday := (newYearsDay.Weekday()) % 7
	// 週番号の計算
	// 指定した日の日曜日の日付けを取得
	sundayFirstMonth := newYearsDay.AddDate(0, 0, int(-newWeekday))
	//　1/1の週初めから経過した時間を求める
	week := specifiedDate.Sub(sundayFirstMonth)
	// 時間から週番号を取得する
	weeklyNumber := math.Floor((week.Hours() / 24) / 7)
	return int(weeklyNumber) + 1
}
