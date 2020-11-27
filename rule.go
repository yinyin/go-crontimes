package crontimes

import (
	"encoding/binary"
)

// CronRule contains parsed cron rule expression.
type CronRule struct {
	minuteValuePoints         uint64
	hourValuePoints           uint32
	dayValuePoints            uint32
	lastDayOfMonth            bool
	nearestWeekdayValuePoints uint32
	monthValuePoints          uint32
	dayOfWeekValuePoints      uint32
	lastDayOfWeekValuePoints  uint32
}

func (r *CronRule) setMinutePattern(patternText string) (err error) {
	elem, err := parseElementText(patternText, 0, 59, 59, false)
	if nil != err {
		return
	}
	if elem.LastValues != 0 {
		return newErrCronExpression(ErrLastNotSupportInMinuteRule, 0)
	}
	if elem.WeekdayValues != 0 {
		return newErrCronExpression(ErrWeekdayNotSupportInMinuteRule, 0)
	}
	if elem.ValuePoints == 0 {
		r.minuteValuePoints = 0x1
	} else {
		r.minuteValuePoints = elem.ValuePoints
	}
	return
}

func (r *CronRule) setHourPattern(patternText string) (err error) {
	elem, err := parseElementText(patternText, 0, 23, 23, false)
	if nil != err {
		return
	}
	if elem.LastValues != 0 {
		return newErrCronExpression(ErrLastNotSupportInHourRule, 0)
	}
	if elem.WeekdayValues != 0 {
		return newErrCronExpression(ErrWeekdayNotSupportInHourRule, 0)
	}
	if elem.ValuePoints == 0 {
		r.hourValuePoints = 0x1
	} else {
		r.hourValuePoints = uint32(elem.ValuePoints)
	}
	return
}

func (r *CronRule) setDayPattern(patternText string) (err error) {
	elem, err := parseElementText(patternText, 1, 31, 31, true)
	if nil != err {
		return
	}
	switch elem.LastValues {
	case 0:
		r.lastDayOfMonth = false
	case 1:
		r.lastDayOfMonth = true
	default:
		return newErrCronExpression(ErrLastDayOfMonthSyntax, 0)
	}
	r.nearestWeekdayValuePoints = uint32(elem.WeekdayValues)
	if elem.ValuePoints == 0 {
		r.dayValuePoints = 0xFFFFFFFE
	} else {
		r.dayValuePoints = uint32(elem.ValuePoints)
	}
	return
}

func (r *CronRule) setMonthPattern(patternText string) (err error) {
	elem, err := parseElementText(patternText, 1, 12, 12, false)
	if nil != err {
		return
	}
	if elem.LastValues != 0 {
		return newErrCronExpression(ErrLastNotSupportInMonthRule, 0)
	}
	if elem.WeekdayValues != 0 {
		return newErrCronExpression(ErrWeekdayNotSupportInMonthRule, 0)
	}
	if elem.ValuePoints == 0 {
		r.monthValuePoints = 0x1FFE
	} else {
		r.monthValuePoints = uint32(elem.ValuePoints)
	}
	return
}

func (r *CronRule) setDayOfWeekPattern(patternText string) (err error) {
	elem, err := parseElementText(patternText, 0, 6, 7, false)
	if nil != err {
		return
	}
	if elem.WeekdayValues != 0 {
		return newErrCronExpression(ErrWeekdayNotSupportInDayOfWeekRule, 0)
	}
	if aux := uint32(elem.LastValues); aux == 0 {
		r.lastDayOfWeekValuePoints = 0
	} else {
		r.lastDayOfWeekValuePoints = (aux | ((aux >> 7) & 1)) & 0x7F
	}
	if aux := uint32(elem.ValuePoints); aux == 0 {
		r.dayOfWeekValuePoints = 0x7F
	} else {
		r.dayOfWeekValuePoints = (aux | ((aux >> 7) & 1)) & 0x7F
	}
	return
}

// SetRule parse and assign given cron patterns.
func (r *CronRule) SetRule(minutePattern, hourPattern, dayPattern, monthPattern, dayOfWeekPattern string) (err error) {
	if err = r.setMinutePattern(minutePattern); nil != err {
		return
	}
	if err = r.setHourPattern(hourPattern); nil != err {
		return
	}
	if err = r.setDayPattern(dayPattern); nil != err {
		return
	}
	if err = r.setMonthPattern(monthPattern); nil != err {
		return
	}
	if err = r.setDayOfWeekPattern(dayOfWeekPattern); nil != err {
		return
	}
	return
}

// MarshalBinary implements encoding.BinaryMarshaler interface.
func (r *CronRule) MarshalBinary() (data []byte, err error) {
	var flags uint32
	if r.lastDayOfMonth {
		flags = 0x1
	}
	data = make([]byte, 36)
	binary.LittleEndian.PutUint64(data[0:], r.minuteValuePoints)
	binary.LittleEndian.PutUint32(data[8:], r.hourValuePoints)
	binary.LittleEndian.PutUint32(data[12:], r.dayValuePoints)
	binary.LittleEndian.PutUint32(data[16:], r.nearestWeekdayValuePoints)
	binary.LittleEndian.PutUint32(data[20:], r.monthValuePoints)
	binary.LittleEndian.PutUint32(data[24:], r.dayOfWeekValuePoints)
	binary.LittleEndian.PutUint32(data[28:], r.lastDayOfWeekValuePoints)
	binary.LittleEndian.PutUint32(data[32:], flags)
	return
}

// UnmarshalBinary implement encoding.BinaryUnmarshaler interface.
func (r *CronRule) UnmarshalBinary(data []byte) (err error) {
	if len(data) != 36 {
		err = ErrBinarySizeMismatch
		return
	}
	r.minuteValuePoints = binary.LittleEndian.Uint64(data[0:])
	r.hourValuePoints = binary.LittleEndian.Uint32(data[8:])
	r.dayValuePoints = binary.LittleEndian.Uint32(data[12:])
	r.nearestWeekdayValuePoints = binary.LittleEndian.Uint32(data[16:])
	r.monthValuePoints = binary.LittleEndian.Uint32(data[20:])
	r.dayOfWeekValuePoints = binary.LittleEndian.Uint32(data[24:])
	r.lastDayOfWeekValuePoints = binary.LittleEndian.Uint32(data[28:])
	flags := binary.LittleEndian.Uint32(data[32:])
	r.lastDayOfMonth = (flags & 0x1) != 0
	return
}
