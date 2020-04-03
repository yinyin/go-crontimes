package crontimes

import (
	"time"
)

const (
	dayOneDayOfWeek  = 1
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
	secondsPerDay    = 24 * secondsPerHour

	secondsPer400Year = (400*365 + 400/4 - 400/100 + 400/400) * secondsPerDay
	secondsPer100Year = (100*365 + 100/4 - 100/100) * secondsPerDay
	secondsPer4Year   = (4*365 + 4/4) * secondsPerDay
	secondsPer1Year   = 365 * secondsPerDay

	secondsInJanuary       = 31 * secondsPerDay
	secondsBeforeMarch     = 28*secondsPerDay + secondsInJanuary
	secondsBeforeLeapMarch = 29*secondsPerDay + secondsInJanuary

	unixEpochToYearOne int64 = (1969*365 + 1969/4 - 1969/100 + 1969/400) * secondsPerDay
	yearOneToUnixEpoch int64 = -unixEpochToYearOne
)

type secondsPerGivenYears struct {
	seconds int64
	years   int
}

var secondsPerYears = [...]secondsPerGivenYears{
	{
		seconds: secondsPer400Year,
		years:   400,
	},
	{
		seconds: secondsPer100Year,
		years:   100,
	},
	{
		seconds: secondsPer4Year,
		years:   4,
	},
	{
		seconds: secondsPer1Year,
		years:   1,
	},
}

func yearFromYearOneSeconds(ts int64) (remainSec int64, year int) {
	year = 1
	remainSec = ts
	for _, secPerYear := range secondsPerYears {
		if remainSec > secPerYear.seconds {
			year += secPerYear.years * int(remainSec/secPerYear.seconds)
			remainSec = remainSec % secPerYear.seconds
		}
	}
	return
}

var secondsBeforeGivenMonthAfterFeburary = [...]int64{
	0,
	31 * secondsPerDay,
	(31 + 30) * secondsPerDay,
	(31 + 30 + 31) * secondsPerDay,
	(31 + 30 + 31 + 30) * secondsPerDay,
	(31 + 30 + 31 + 30 + 31) * secondsPerDay,
	(31 + 30 + 31 + 30 + 31 + 31) * secondsPerDay,
	(31 + 30 + 31 + 30 + 31 + 31 + 30) * secondsPerDay,
	(31 + 30 + 31 + 30 + 31 + 31 + 30 + 31) * secondsPerDay,
	(31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30) * secondsPerDay,
	(31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31) * secondsPerDay,
}

func isLeapYear(year int) bool {
	return ((year % 4) == 0) && (((year % 100) != 0) || ((year % 400) == 0))
}

var daysOfMonthNormalYear = [...]int{
	0,
	31, 28, 31,
	30, 31, 30,
	31, 31, 30,
	31, 30, 31,
}

var daysOfMonthLeapYear = [...]int{
	0,
	31, 29, 31,
	30, 31, 30,
	31, 31, 30,
	31, 30, 31,
}

func numberOfDaysForMonth(month int, inLeapYear bool) int {
	if inLeapYear {
		return daysOfMonthLeapYear[month]
	}
	return daysOfMonthNormalYear[month]
}

func monthFromJanuaryOneSeconds(s int64, year int) (remainSec int64, month int) {
	if s < secondsInJanuary {
		month = 1
		remainSec = s
		return
	}
	inLeapYear := isLeapYear(year)
	var secBeforeMarch int64
	if !inLeapYear {
		secBeforeMarch = secondsBeforeMarch
	} else {
		secBeforeMarch = secondsBeforeLeapMarch
	}
	if s < secBeforeMarch {
		month = 2
		remainSec = s - secondsInJanuary
		return
	}
	month = 3
	remainSec = s - secBeforeMarch
	estMonthOffset := int(remainSec / (31 * secondsPerDay))
	if remainSec < secondsBeforeGivenMonthAfterFeburary[estMonthOffset+1] {
		month += estMonthOffset
		remainSec -= secondsBeforeGivenMonthAfterFeburary[estMonthOffset]
	} else {
		month += (estMonthOffset + 1)
		remainSec -= secondsBeforeGivenMonthAfterFeburary[estMonthOffset+1]
	}
	return
}

func dayOfMonthFromDayOneSeconds(s int64) (remainSec int64, day int) {
	remainSec = s % secondsPerDay
	day = int(s/secondsPerDay) + 1
	return
}

func hoursFromHourZeroSeconds(s int64) (remainSec int64, hour int) {
	remainSec = s % secondsPerHour
	hour = int(s / secondsPerHour)
	return
}

func minutesFromMinuteZeroSeconds(s int64) (remainSec int64, minute int) {
	remainSec = s % secondsPerMinute
	minute = int(s / secondsPerMinute)
	return
}

func dayOfWeekFromYearOneSeconds(ts int64) (dayOfWeek int) {
	dayOfWeek = int((int64(ts/secondsPerDay) + dayOneDayOfWeek) % 7)
	return
}

func factorizeTime(t time.Time) (localYearOneMinuteAlignedSec int64, tzOffset, year, month, day, hour, minute, dayOfWeek int) {
	unixSec := t.Unix()
	_, tzOffset = t.Zone()
	localYearOneSec := unixSec + int64(tzOffset) + unixEpochToYearOne
	remainSec, year := yearFromYearOneSeconds(localYearOneSec)
	remainSec, month = monthFromJanuaryOneSeconds(remainSec, year)
	remainSec, day = dayOfMonthFromDayOneSeconds(remainSec)
	remainSec, hour = hoursFromHourZeroSeconds(remainSec)
	remainSec, minute = minutesFromMinuteZeroSeconds(remainSec)
	dayOfWeek = dayOfWeekFromYearOneSeconds(localYearOneSec)
	localYearOneMinuteAlignedSec = localYearOneSec - remainSec
	return
}

// CronTimes is an iterator for given cron expression and time range.
type CronTimes struct {
	CronRule

	location         *time.Location
	localYearOneSec  int64
	tzOffset         int
	year, month, day int
	hour, minute     int
	dayOfWeek        int
	inLeapYear       bool

	boundaryUnixSec         int64
	boundaryLocalYearOneSec int64
}

// SetRange set iterator range from tStartIncl to tEndExcl with given location.
func (c *CronTimes) SetRange(tStartIncl, tEndExcl time.Time, loc *time.Location) {
	if loc == nil {
		loc = time.UTC
	}
	s := tStartIncl.Add(time.Second * 59).In(loc)
	c.location = loc
	c.localYearOneSec, c.tzOffset, c.year, c.month, c.day, c.hour, c.minute, c.dayOfWeek = factorizeTime(s)
	c.inLeapYear = isLeapYear(c.year)
	c.boundaryUnixSec = tEndExcl.Unix()
	c.recomputeBoundaryLocalYearOneSec()
}

func (c *CronTimes) recomputeBoundaryLocalYearOneSec() {
	c.boundaryLocalYearOneSec = c.boundaryUnixSec + int64(c.tzOffset) + unixEpochToYearOne
}

func (c *CronTimes) workingUnix() (u int64) {
	u = c.localYearOneSec - int64(c.tzOffset) + yearOneToUnixEpoch
	return
}

func (c *CronTimes) workingTime() (t time.Time) {
	u := c.workingUnix()
	t = time.Unix(u, 0).In(c.location)
	return
}

func (c *CronTimes) increaseMonth() {
	if c.month == 12 {
		c.year++
		c.month = 1
		c.inLeapYear = isLeapYear(c.year)
		return
	}
	c.month++
}

func (c *CronTimes) increaseDay() {
	c.dayOfWeek = (c.dayOfWeek + 1) % 7
	if c.day == numberOfDaysForMonth(c.month, c.inLeapYear) {
		c.increaseMonth()
		c.day = 1
		return
	}
	c.day++
}

func (c *CronTimes) increaseHour() {
	if c.hour == 23 {
		c.increaseDay()
		c.hour = 0
		return
	}
	c.hour++
}

func (c *CronTimes) increaseMinute() {
	if c.minute == 59 {
		c.increaseHour()
		c.minute = 0
		return
	}
	c.minute++
}

func (c *CronTimes) moveToNextMinute() {
	c.increaseMinute()
	c.localYearOneSec += 60
}

func (c *CronTimes) moveToNextHour() {
	mins := 60 - c.minute
	c.increaseHour()
	c.minute = 0
	c.localYearOneSec += int64(mins * 60)
}

func (c *CronTimes) moveToNextDay() {
	mins := 60 - c.minute
	hours := 23 - c.hour
	c.increaseDay()
	c.minute = 0
	c.hour = 0
	c.localYearOneSec += int64((hours * 3600) + (mins * 60))
}

// isMatchingNearestWeekday check if day is matching to nearest weekday.
// Includes rule like: 3W
func (c *CronTimes) isMatchingNearestWeekday() bool {
	if (c.nearestWeekdayValuePoints == 0) || (c.dayOfWeek == 0) || (c.dayOfWeek == 6) {
		return false
	}
	var mask uint32
	switch c.dayOfWeek {
	case 1:
		switch c.day {
		case 1:
			mask = (1 << 1)
		case 3:
			mask = (1 << 1) | (1 << 2) | (1 << 3)
		default:
			mask = (1 << (c.day - 1)) | (1 << c.day)
		}
	case 5:
		n := numberOfDaysForMonth(c.month, c.inLeapYear)
		switch c.day {
		case n:
			mask = (1 << n)
		case n - 2:
			mask = (1 << (n - 2)) | (1 << (n - 1)) | (1 << n)
		default:
			mask = (1 << c.day) | (1 << (c.day + 1))
		}
	default:
		mask = 1 << c.day
	}
	return ((c.nearestWeekdayValuePoints & mask) != 0)
}

// isMatchingLastDayOfMonth check if day is matching to given last day of month day rule.
// Includes rule like: L
func (c *CronTimes) isMatchingLastDayOfMonth() bool {
	return (c.lastDayOfMonth && (numberOfDaysForMonth(c.month, c.inLeapYear) == c.day))
}

// isMatchingDay check if day is matching to given day rule.
// Include rules like: *, 1-5, */6
func (c *CronTimes) isMatchingDay() bool {
	return (((1 << c.day) & c.dayValuePoints) != 0)
}

// isMatchingLastDayOfWeek check if day of week is matching to last day of week rule.
// Includes rule like: 1L
func (c *CronTimes) isMatchingLastDayOfWeek() bool {
	return (c.lastDayOfWeekValuePoints != 0) &&
		(((1 << c.dayOfWeek) & c.lastDayOfWeekValuePoints) != 0) &&
		((c.day + 7) > numberOfDaysForMonth(c.month, c.inLeapYear))
}

// isMatchingMonth check if month is matched.
// Include rules like: *, */3, 2,5,8,11, 7-10
func (c *CronTimes) isMatchingMonth() bool {
	return (((1 << c.month) & c.monthValuePoints) != 0)
}

// isMatchingDayOfWeek check if day of week is matched.
// Include rules like: *, 1-3, */2
func (c *CronTimes) isMatchingDayOfWeek() bool {
	return (((1 << c.dayOfWeek) & c.dayOfWeekValuePoints) != 0)
}

func (c *CronTimes) isMatchingHour() bool {
	return ((1 << c.hour) & c.hourValuePoints) != 0
}

func (c *CronTimes) isMatchingMinute() bool {
	return ((1 << c.minute) & c.minuteValuePoints) != 0
}

// updateDST check if hour offset is changed.
// Local time will be update if hour offset is changed.
// Only invoke this check on 0 minute.
func (c *CronTimes) updateDST() {
	w := c.workingTime()
	if _, tzOffset := w.Zone(); tzOffset == c.tzOffset {
		return
	}
	workingTick := c.localYearOneSec
	c.localYearOneSec, c.tzOffset, c.year, c.month, c.day, c.hour, c.minute, c.dayOfWeek = factorizeTime(w)
	c.recomputeBoundaryLocalYearOneSec()
	for workingTick > c.localYearOneSec {
		c.moveToNextHour()
	}
}

// iterationCompleted check if working time-stamp exceeded boundary time.
func (c *CronTimes) iterationCompleted() bool {
	return (c.localYearOneSec >= c.boundaryLocalYearOneSec)
}

type matchStatus int

const (
	matchNotFound matchStatus = iota
	matchFound
	matchExceedBoundary
)

// internalNextUnix search for next match time instant.
func (c *CronTimes) seekNextMatch() matchStatus {
	if c.iterationCompleted() {
		return matchExceedBoundary
	}
	for !(c.isMatchingMonth() &&
		(c.isMatchingDay() || c.isMatchingLastDayOfMonth() || c.isMatchingNearestWeekday()) &&
		(c.isMatchingDayOfWeek() || c.isMatchingLastDayOfWeek())) {
		c.moveToNextDay()
		c.updateDST()
		if c.iterationCompleted() {
			return matchExceedBoundary
		}
	}
	for !c.isMatchingHour() {
		workingDay := c.day
		c.moveToNextHour()
		c.updateDST()
		if workingDay != c.day {
			return matchNotFound
		}
		if c.iterationCompleted() {
			return matchExceedBoundary
		}
	}
	for !c.isMatchingMinute() {
		c.moveToNextMinute()
		if c.minute == 0 {
			c.updateDST()
			return matchNotFound
		}
		if c.iterationCompleted() {
			return matchExceedBoundary
		}
	}
	return matchFound
}

// NextUnix return next matched Unix timestamp in seconds.
// Return 0 if not matches in given range.
func (c *CronTimes) NextUnix() (result int64) {
	m := c.seekNextMatch()
	for m == matchNotFound {
		m = c.seekNextMatch()
	}
	if m == matchExceedBoundary {
		return
	}
	result = c.workingUnix()
	c.moveToNextMinute()
	if c.minute == 0 {
		c.updateDST()
	}
	return
}

// NextTime return next matched time.Time.
// Result will be zero if not matches in given range.
func (c *CronTimes) NextTime() (result time.Time) {
	if t := c.NextUnix(); t != 0 {
		result = time.Unix(t, 0).In(c.location)
	}
	return
}
