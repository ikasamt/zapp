package zapp

import "time"

type CalendarSupport struct {
	Now time.Time
}

// 先月初日
func (x CalendarSupport) FirstDayOfPrevMonth() time.Time {
	t := x.Now
	firstDayOfThisMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	firstDayOfPrevMonth := firstDayOfThisMonth.AddDate(0, -1, 0)
	return firstDayOfPrevMonth
}

// 先月の「月」の数字。例: 1月なら1、12月なら12
func (x CalendarSupport) PrevMonth() string {
	return x.FirstDayOfPrevMonth().Format(`1`)
}

// 先月の「月」の数字。例: 2018年1月なら201801、2018年12月なら201812
func (x CalendarSupport) PrevYearMonth() string {
	return x.FirstDayOfPrevMonth().Format(`200601`)
}

// 先月の「月」の数字。例: 1月なら1、12月なら12
func (x CalendarSupport) ThisMonth() string {
	return x.Now.Format(`1`)
}

// 先月の「月」の数字。例: 1月なら1、12月なら12
func (x CalendarSupport) ThisYearMonth() string {
	return x.Now.Format(`200601`)
}
