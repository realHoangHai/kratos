package util

import "time"

// GetYesterdayRangeTime Get interval time - yesterday
func GetYesterdayRangeTime() (time.Time, time.Time) {
	now := time.Now().AddDate(0, 0, -1)
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return startDate, endDate
}

// GetTodayRangeTime Get interval time - today
func GetTodayRangeTime() (time.Time, time.Time) {
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return startDate, endDate
}

// GetLastMonthRangeTime Get interval time - last month
func GetLastMonthRangeTime() (time.Time, time.Time) {
	now := time.Now().AddDate(0, -1, 0)
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetCurrentMonthRangeTime Get interval time - this month
func GetCurrentMonthRangeTime() (time.Time, time.Time) {
	now := time.Now()
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetCurrentYearRangeTime Get interval time - this year
func GetCurrentYearRangeTime() (time.Time, time.Time) {
	now := time.Now()
	firstDay := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(1, 0, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetLastYearRangeTime Get interval time - last year
func GetLastYearRangeTime() (time.Time, time.Time) {
	now := time.Now().AddDate(-1, 0, 0)
	firstDay := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(1, 0, -1)
	endDate := time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 23, 59, 59, 0, now.Location())
	return firstDay, endDate
}

// GetTodayRangeDateString Get the date string within the interval (for example: 2023-05-23 - 2023-05-23) - today
func GetTodayRangeDateString() (string, string) {
	firstDay, lastDay := GetTodayRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetYesterdayRangeDateString Get the date string within the interval (for example: 2023-05-23 - 2023-05-23) - yesterday
func GetYesterdayRangeDateString() (string, string) {
	firstDay, lastDay := GetYesterdayRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetCurrentMonthRangeDateString Get the date string within the interval (for example: 2023-05-23 - 2023-05-23) - this month
func GetCurrentMonthRangeDateString() (string, string) {
	firstDay, lastDay := GetCurrentMonthRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetLastMonthRangeDateString Get the date string within the interval (for example: 2023-05-23 - 2023-05-23) - last month
func GetLastMonthRangeDateString() (string, string) {
	firstDay, lastDay := GetLastMonthRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetCurrentYearRangeDateString Get the date string within the interval (for example: 2023-05-23 - 2023-05-23) - this year
func GetCurrentYearRangeDateString() (string, string) {
	firstDay, lastDay := GetCurrentYearRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetLastYearRangeDateString Get the date string within the interval (for example: 2023-05-23 - 2023-05-23) - last year
func GetLastYearRangeDateString() (string, string) {
	firstDay, lastDay := GetLastYearRangeTime()

	startDate := firstDay.Format(DateLayout)
	endDate := lastDay.Format(DateLayout)

	return startDate, endDate
}

// GetYesterdayRangeTimeString Get the time string within the interval (for example: 2023-05-23 00:00:00 2023-05-23 23:59:59) - yesterday
func GetYesterdayRangeTimeString() (string, string) {
	firstDay, lastDay := GetYesterdayRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetTodayRangeTimeString Get the time string within the interval (for example: 2023-05-23 00:00:00 2023-05-23 23:59:59) - today
func GetTodayRangeTimeString() (string, string) {
	firstDay, lastDay := GetTodayRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetLastMonthRangeTimeString Get the time string within the interval (for example: 2023-05-23 00:00:00 2023-05-23 23:59:59) - last month
func GetLastMonthRangeTimeString() (string, string) {
	firstDay, lastDay := GetLastMonthRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetCurrentMonthRangeTimeString Get the time string within the interval (for example: 2023-05-23 00:00:00 2023-05-23 23:59:59) - this month
func GetCurrentMonthRangeTimeString() (string, string) {
	firstDay, lastDay := GetCurrentMonthRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetLastYearRangeTimeString Get the time string within the interval (for example: 2023-05-23 00:00:00 2023-05-23 23:59:59) - last year
func GetLastYearRangeTimeString() (string, string) {
	firstDay, lastDay := GetLastYearRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}

// GetCurrentYearRangeTimeString Get the time string within the interval (for example: 2023-05-23 00:00:00 2023-05-23 23:59:59) - this year
func GetCurrentYearRangeTimeString() (string, string) {
	firstDay, lastDay := GetCurrentYearRangeTime()

	startDate := firstDay.Format(TimeLayout)
	endDate := lastDay.Format(TimeLayout)

	return startDate, endDate
}
