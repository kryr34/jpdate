package main

import (
	"os"
	"fmt"
	"time"
	"regexp"
	"strconv"
)

func japanHour(hour, daySplit int) string {
	daySplit %= 24
	if hour < daySplit {
		return strconv.Itoa(hour+24)
	}
	return "%H"
}

func japanWeekday(weekday time.Weekday) string {
	strArr := [7]string{"日", "月", "火", "水", "木", "金", "土"}
	return strArr[weekday]
}

func main() {
	current			:= time.Now()
	formatString	:= "%Y年%m月%_d日 (%Ja) %JH:%M:%S %Z"
	
	argLength := len(os.Args[1:])
	if argLength >= 1 {
		formatString = os.Args[1]
	}

	jh := japanHour(current.Hour(), 30)
	ja := japanWeekday(current.Weekday())
	
	formatString = regexp.MustCompile(`%JH`).
		ReplaceAllString(formatString, jh)
	formatString = regexp.MustCompile(`%Ja`).
		ReplaceAllString(formatString, ja)
	formatString = regexp.MustCompile(`%JA`).
		ReplaceAllString(formatString, ja+"曜日")
	fmt.Println(formatString)
}
