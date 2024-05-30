package nepdatego

import "time"

func ParseDate(input, pattern string) map[rune]int {
	time_and_date := []rune{}
	pat := map[rune]int{}
	next := false
	for _, part := range pattern {
		if part == '%' {
			next = true
		} else if next {
			time_and_date = append(time_and_date, part)
			next = false
		}
	}

	num, cur := 0, 0
	for _, v := range input {
		if v < '0' || v > '9' {
			if num > 0 {
				pat[time_and_date[cur]] = num
				num, cur = 0, cur+1
			}
		} else {
			num = num*10 + int(v-'0')
		}
	}
	if cur < len(time_and_date) {
		pat[time_and_date[cur]] = num
	}
	return pat
}

func GetDate(input, pattern string, time_zone *time.Location) time.Time {
	date := ParseDate(input, pattern)
	return time.Date(date['Y'], time.Month(date['M']), date['D'], date['h'], date['m'], date['s'], date['n'], time_zone).AddDate(57, 8, 16)
}
