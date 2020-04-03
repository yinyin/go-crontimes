package crontimes

import (
	"testing"
	"time"
)

func TestFactorizeTime_TPE20200325010339(t *testing.T) {
	p0 := time.Unix(1585069419, 0) // 2020-03-25 01:03:39 (Wed)
	tz, err := time.LoadLocation("Asia/Taipei")
	if nil != err {
		t.Fatalf("cannot load time zone: %v", err)
	}
	p1 := p0.In(tz)
	y1MinuteAlignedSec, tzOffset, year, month, day, hour, minute, dayOfWeek := factorizeTime(p1)
	if y1MinuteAlignedSec != 63720694980 {
		t.Errorf("unexpect y1MinuteAlignedSec: %d", y1MinuteAlignedSec)
	}
	if tzOffset != 8*3600 {
		t.Errorf("unexpect tzOffset: %d", tzOffset)
	}
	if year != 2020 {
		t.Errorf("unexpect year: %d", year)
	}
	if month != 3 {
		t.Errorf("unexpect month: %d", month)
	}
	if day != 25 {
		t.Errorf("unexpect day: %d", day)
	}
	if hour != 1 {
		t.Errorf("unexpect hour: %d", hour)
	}
	if minute != 3 {
		t.Errorf("unexpect minute: %d", minute)
	}
	if dayOfWeek != 3 {
		t.Errorf("unexpect dayOfWeek: %d", dayOfWeek)
	}
}

func TestCronTimesSetRangeMoveOperations1(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Taipei")
	if nil != err {
		t.Fatalf("cannot load time zone: %v", err)
	}
	t0 := time.Date(1999, 12, 29, 8, 15, 30, 199, tz)
	t1 := time.Date(2000, 3, 10, 8, 15, 30, 199, tz)
	cronTimes := CronTimes{}
	cronTimes.SetRange(t0, t1, tz)
	tAligned := time.Date(1999, 12, 29, 8, 16, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(tAligned) {
		t.Errorf("unexpect working time (aligned): %v", w)
		return
	}
	cronTimes.moveToNextMinute()
	t1MinNext := time.Date(1999, 12, 29, 8, 17, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(t1MinNext) {
		t.Errorf("unexpect working time (move-1min): %v", w)
		return
	}
	cronTimes.moveToNextHour()
	t1HrNext := time.Date(1999, 12, 29, 9, 0, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(t1HrNext) {
		t.Errorf("unexpect working time (move-1hour): %v", w)
		return
	}
	cronTimes.moveToNextDay()
	t1DayNext := time.Date(1999, 12, 30, 0, 0, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(t1DayNext) {
		t.Errorf("unexpect working time (move-1day): %v", w)
		return
	}
}

func TestCronTimesSetRangeMoveOperations2(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Taipei")
	if nil != err {
		t.Fatalf("cannot load time zone: %v", err)
	}
	t0 := time.Date(1999, 12, 30, 22, 58, 30, 199, tz)
	t1 := time.Date(2000, 3, 10, 8, 15, 30, 199, tz)
	cronTimes := CronTimes{}
	cronTimes.SetRange(t0, t1, tz)
	tAligned := time.Date(1999, 12, 30, 22, 59, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(tAligned) {
		t.Errorf("unexpect working time (aligned): %v", w)
		return
	}
	cronTimes.moveToNextMinute()
	t1MinNext := time.Date(1999, 12, 30, 23, 00, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(t1MinNext) {
		t.Errorf("unexpect working time (move-1min): %v", w)
		return
	}
	cronTimes.moveToNextHour()
	t1HrNext := time.Date(1999, 12, 31, 0, 0, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(t1HrNext) {
		t.Errorf("unexpect working time (move-1hour): %v", w)
		return
	}
	cronTimes.moveToNextDay()
	t1DayNext := time.Date(2000, 1, 1, 0, 0, 0, 0, tz)
	if w := cronTimes.workingTime(); !w.Equal(t1DayNext) {
		t.Errorf("unexpect working time (move-1day): %v", w)
		return
	}
}

func TestCronTimesIterator1(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Taipei")
	if nil != err {
		t.Fatalf("cannot load time zone: %v", err)
	}
	t0 := time.Date(1999, 12, 29, 8, 15, 30, 199, tz)
	t1 := time.Date(2000, 3, 10, 8, 15, 30, 199, tz)
	cronTimes := CronTimes{}
	if err = cronTimes.SetRule("*", "*", "*", "*", "*"); nil != err {
		t.Errorf("cannot set rule: %v", err)
		return
	}
	cronTimes.SetRange(t0, t1, tz)
	tCmp := time.Date(1999, 12, 29, 8, 16, 0, 0, tz)
	cycleCount := 0
	tHave := cronTimes.NextTime()
	for !tHave.IsZero() {
		if !tHave.Equal(tCmp) {
			t.Errorf("unexpect value (%v vs. %v; cycle %d)", tHave, tCmp, cycleCount)
			return
		}
		tCmp = tCmp.Add(time.Minute)
		tHave = cronTimes.NextTime()
		cycleCount++
	}
	b0 := time.Date(1999, 12, 29, 8, 16, 0, 0, tz)
	b1 := time.Date(2000, 3, 10, 8, 15, 0, 0, tz)
	expCount := int((b1.Unix()-b0.Unix())/60) + 1
	if expCount != cycleCount {
		t.Errorf("unexpect cycle count: %d vs. %d", expCount, cycleCount)
	}
}

func TestCronTimesIterator2DST(t *testing.T) {
	tz, err := time.LoadLocation("America/Chicago")
	if nil != err {
		t.Fatalf("cannot load time zone: %v", err)
	}
	t0 := time.Date(2020, 3, 7, 2, 29, 30, 199, tz)
	t1 := time.Date(2020, 3, 9, 22, 15, 30, 199, tz)
	cronTimes := CronTimes{}
	if err = cronTimes.SetRule("*", "*", "*", "*", "*"); nil != err {
		t.Errorf("cannot set rule: %v", err)
		return
	}
	cronTimes.SetRange(t0, t1, tz)
	tCmp := time.Date(2020, 3, 7, 2, 30, 0, 0, tz)
	cycleCount := 0
	tHave := cronTimes.NextTime()
	for !tHave.IsZero() {
		if !tHave.Equal(tCmp) {
			t.Errorf("unexpect value (%v vs. %v; cycle %d)", tHave, tCmp, cycleCount)
			return
		}
		tCmp = tCmp.Add(time.Minute)
		tHave = cronTimes.NextTime()
		cycleCount++
	}
	b0 := time.Date(2020, 3, 7, 2, 30, 0, 0, tz)
	b1 := time.Date(2020, 3, 9, 22, 15, 0, 0, tz)
	expCount := int((b1.Unix()-b0.Unix())/60) + 1
	if expCount != cycleCount {
		t.Errorf("unexpect cycle count: %d vs. %d", expCount, cycleCount)
	}
}

func TestCronTimesIterator3DST(t *testing.T) {
	tz, err := time.LoadLocation("America/Chicago")
	if nil != err {
		t.Fatalf("cannot load time zone: %v", err)
	}
	t0 := time.Date(2020, 10, 29, 2, 29, 30, 199, tz)
	t1 := time.Date(2020, 11, 3, 22, 15, 30, 199, tz)
	cronTimes := CronTimes{}
	if err = cronTimes.SetRule("*", "*", "*", "*", "*"); nil != err {
		t.Errorf("cannot set rule: %v", err)
		return
	}
	cronTimes.SetRange(t0, t1, tz)
	tCmp := time.Date(2020, 10, 29, 2, 30, 0, 0, tz)
	tDSTSwitch := time.Date(2020, 11, 1, 7, 0, 0, 0, time.UTC).In(tz)
	cycleCount := 0
	tHave := cronTimes.NextTime()
	for !tHave.IsZero() {
		if !tHave.Equal(tCmp) {
			t.Errorf("unexpect value (%v vs. %v; cycle %d)", tHave, tCmp, cycleCount)
			return
		}
		tCmp = tCmp.Add(time.Minute)
		if tDSTSwitch.Equal(tCmp) {
			tCmp = tCmp.Add(time.Hour)
		}
		tHave = cronTimes.NextTime()
		cycleCount++
	}
	b0 := time.Date(2020, 10, 29, 2, 30, 0, 0, tz)
	b1 := time.Date(2020, 11, 3, 22, 15, 0, 0, tz)
	expCount := int((b1.Unix()-b0.Unix()-3600)/60) + 1
	if expCount != cycleCount {
		t.Errorf("unexpect cycle count: %d vs. %d", expCount, cycleCount)
	}
}
