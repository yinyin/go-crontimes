package crontimes

import (
	"testing"
)

func TestCronRuleSetMinutePattern01(t *testing.T) {
	var r CronRule
	if err := r.setMinutePattern("*"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.minuteValuePoints != 0xFFFFFFFFFFFFFFF {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.minuteValuePoints)
		return
	}
	if err := r.setMinutePattern("*/2"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.minuteValuePoints != 0x555555555555555 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.minuteValuePoints)
		return
	}
	if err := r.setMinutePattern("0,59"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.minuteValuePoints != 0x800000000000001 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.minuteValuePoints)
		return
	}
	if err := r.setMinutePattern("2-10"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.minuteValuePoints != 0x07FC {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.minuteValuePoints)
		return
	}
	if err := r.setMinutePattern("2-10/3"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.minuteValuePoints != 0x0124 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.minuteValuePoints)
		return
	}
	if err := r.setMinutePattern(""); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.minuteValuePoints != 0x01 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.minuteValuePoints)
		return
	}
}

func TestCronRuleSetMinutePattern02(t *testing.T) {
	var r CronRule
	if err := r.setMinutePattern("L"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrLastNotSupportInMinuteRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
	if err := r.setMinutePattern("2L"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrLastNotSupportInMinuteRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
	if err := r.setMinutePattern("1W"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrWeekdayNotSupportInMinuteRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
}

func TestCronRuleSetHourPattern01(t *testing.T) {
	var r CronRule
	if err := r.setHourPattern("*"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.hourValuePoints != 0xFFFFFF {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.hourValuePoints)
		return
	}
	if err := r.setHourPattern("*/2"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.hourValuePoints != 0x555555 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.hourValuePoints)
		return
	}
	if err := r.setHourPattern("0,23"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.hourValuePoints != 0x800001 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.hourValuePoints)
		return
	}
	if err := r.setHourPattern("2-10"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.hourValuePoints != 0x07FC {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.hourValuePoints)
		return
	}
	if err := r.setHourPattern("2-10/3"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.hourValuePoints != 0x0124 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.hourValuePoints)
		return
	}
	if err := r.setHourPattern(""); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.hourValuePoints != 0x01 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.hourValuePoints)
		return
	}
}

func TestCronRuleSetHourPattern02(t *testing.T) {
	var r CronRule
	if err := r.setHourPattern("L"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrLastNotSupportInHourRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
	if err := r.setHourPattern("2L"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrLastNotSupportInHourRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
	if err := r.setHourPattern("1W"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrWeekdayNotSupportInHourRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
}

func TestCronRuleSetDayPattern01(t *testing.T) {
	var r CronRule
	if err := r.setDayPattern("*"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayValuePoints != 0xFFFFFFFE {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayValuePoints)
		return
	}
	if err := r.setDayPattern("*/2"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayValuePoints != 0xAAAAAAAA {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayValuePoints)
		return
	}
	if err := r.setDayPattern("1,23"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayValuePoints != 0x800002 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayValuePoints)
		return
	}
	if err := r.setDayPattern("2-10"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayValuePoints != 0x07FC {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayValuePoints)
		return
	}
	if err := r.setDayPattern("2-10/3"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayValuePoints != 0x0124 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayValuePoints)
		return
	}
	if err := r.setDayPattern(""); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayValuePoints != 0xFFFFFFFE {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayValuePoints)
		return
	}
	if err := r.setDayPattern("L"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if !r.lastDayOfMonth {
		t.Errorf("unexpect result: %#v; %v", &r, r.lastDayOfMonth)
		return
	}
	if err := r.setDayPattern("3W"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.nearestWeekdayValuePoints != 0x08 {
		t.Errorf("unexpect result: %#v; %v", &r, r.nearestWeekdayValuePoints)
		return
	}
	if err := r.setDayPattern("3W,5W,7"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.nearestWeekdayValuePoints != 0x28 {
		t.Errorf("unexpect result: %#v; %v", &r, r.nearestWeekdayValuePoints)
		return
	}
	if r.dayValuePoints != 0x80 {
		t.Errorf("unexpect result: %#v; %v", &r, r.dayValuePoints)
		return
	}
}

func TestCronRuleSetDayPattern02(t *testing.T) {
	var r CronRule
	if err := r.setDayPattern("2L"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrLastDayOfMonthSyntax) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
	if err := r.setDayPattern("W"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrValueOutOfValidRange) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
}

func TestCronRuleSetMonthPattern01(t *testing.T) {
	var r CronRule
	if err := r.setMonthPattern("*"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.monthValuePoints != 0x1FFE {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.monthValuePoints)
		return
	}
	if err := r.setMonthPattern("*/2"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.monthValuePoints != 0x0AAA {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.monthValuePoints)
		return
	}
	if err := r.setMonthPattern("1,12"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.monthValuePoints != 0x1002 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.monthValuePoints)
		return
	}
	if err := r.setMonthPattern("2-10"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.monthValuePoints != 0x07FC {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.monthValuePoints)
		return
	}
	if err := r.setMonthPattern("2-10/3"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.monthValuePoints != 0x0124 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.monthValuePoints)
		return
	}
	if err := r.setMonthPattern(""); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.monthValuePoints != 0x01FFE {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.monthValuePoints)
		return
	}
}

func TestCronRuleSetMonthPattern02(t *testing.T) {
	var r CronRule
	if err := r.setMonthPattern("L"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrValueOutOfValidRange) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
	if err := r.setMonthPattern("2L"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrLastNotSupportInMonthRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
	if err := r.setMonthPattern("1W"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrWeekdayNotSupportInMonthRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
}

func TestCronRuleSetDayOfWeekPattern01(t *testing.T) {
	var r CronRule
	if err := r.setDayOfWeekPattern("*"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x7F {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern("*/2"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x55 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern("1,7"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x03 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern("1,5"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x22 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern("0,1,5"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x23 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern("2-7"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x7D {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern("2-7/3"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x24 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern(""); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.dayOfWeekValuePoints != 0x7F {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if err := r.setDayOfWeekPattern("2L"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	}
	if r.lastDayOfWeekValuePoints != 0x04 {
		t.Errorf("unexpect result: %#v; %v", &r, r.lastDayOfWeekValuePoints)
		return
	}
}

func TestCronRuleSetDayOfWeekPattern02(t *testing.T) {
	var r CronRule
	if err := r.setDayOfWeekPattern("1W"); nil == err {
		t.Errorf("unexpect result: %#v; %v", &r, err)
		return
	} else if e, ok := err.(*ErrCronExpression); (!ok) || (e.ErrType != ErrWeekdayNotSupportInDayOfWeekRule) {
		t.Errorf("unexpect type of error: %v", err)
		return
	}
}

func checkCronRuleSetRule01(t *testing.T, r *CronRule) {
	if r.minuteValuePoints != 0x02 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.minuteValuePoints)
		return
	}
	if r.hourValuePoints != 0x04 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.hourValuePoints)
		return
	}
	if r.dayValuePoints != 0x80000000 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayValuePoints)
		return
	}
	if r.lastDayOfMonth {
		t.Errorf("unexpect result: %#v; %v", &r, r.lastDayOfMonth)
		return
	}
	if r.nearestWeekdayValuePoints != 0 {
		t.Errorf("unexpect result: %#v; %v", &r, r.nearestWeekdayValuePoints)
		return
	}
	if r.monthValuePoints != 0x1000 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.monthValuePoints)
		return
	}
	if r.dayOfWeekValuePoints != 0x20 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
	if r.lastDayOfWeekValuePoints != 0x00 {
		t.Errorf("unexpect result: %#v; 0x%X", &r, r.dayOfWeekValuePoints)
		return
	}
}

func TestCronRuleSetRule01(t *testing.T) {
	var r0, r1 CronRule
	if err := r0.SetRule("1", "2", "31", "12", "5"); nil != err {
		t.Errorf("unexpect result: %#v; %v", &r0, err)
		return
	}
	checkCronRuleSetRule01(t, &r0)
	d, err := r0.MarshalBinary()
	if nil != err {
		t.Errorf("unexpect marshal result: %#v; %v", &r0, err)
		return
	}
	if err = r1.UnmarshalBinary(d); nil != err {
		t.Errorf("unexpect unmarshal result: %#v; %v", &r1, err)
		return
	}
	checkCronRuleSetRule01(t, &r1)
}
