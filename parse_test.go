package crontimes

import (
	"testing"
)

func TestParseElementText01Empty(t *testing.T) {
	r, err := parseElementText("", 0, 60, 60, false)
	if nil != err {
		t.Errorf("unexpected error: %v", err)
	}
	if nil == r {
		t.Error("unexpected nil result")
	} else if !r.isEmpty() {
		t.Error("expecting empty result")
	}
}

func TestParseElementText01UnexpectChar(t *testing.T) {
	if _, err := parseElementText("a", 0, 60, 60, false); nil == err {
		t.Error("expect error")
	} else if e, ok := err.(*ErrCronExpression); !ok {
		t.Errorf("unexpected error type: %#v", err)
	} else if (e.ErrType != ErrSyntaxUnexpectCharacter) || (e.CharIndex != 0) {
		t.Errorf("unexpected error instance: %#v", e)
	}
	if _, err := parseElementText("1,a", 0, 60, 60, false); nil == err {
		t.Error("expect error")
	} else if e, ok := err.(*ErrCronExpression); !ok {
		t.Errorf("unexpected error type: %#v", err)
	} else if (e.ErrType != ErrSyntaxUnexpectCharacter) || (e.CharIndex != 2) {
		t.Errorf("unexpected error instance: %#v", e)
	}
	if _, err := parseElementText("1-a/2", 0, 60, 60, false); nil == err {
		t.Error("expect error")
	} else if e, ok := err.(*ErrCronExpression); !ok {
		t.Errorf("unexpected error type: %#v", err)
	} else if (e.ErrType != ErrSyntaxUnexpectCharacter) || (e.CharIndex != 2) {
		t.Errorf("unexpected error instance: %#v", e)
	}
	if _, err := parseElementText("*/a", 0, 60, 60, false); nil == err {
		t.Error("expect error")
	} else if e, ok := err.(*ErrCronExpression); !ok {
		t.Errorf("unexpected error type: %#v", err)
	} else if (e.ErrType != ErrSyntaxUnexpectCharacter) || (e.CharIndex != 2) {
		t.Errorf("unexpected error instance: %#v", e)
	}
}

func TestParseElementText01OutOfRange(t *testing.T) {
	if _, err := parseElementText("0", 1, 12, 12, false); nil == err {
		t.Error("expect error")
	} else if e, ok := err.(*ErrCronExpression); !ok {
		t.Errorf("unexpected error type: %#v", err)
	} else if (e.ErrType != ErrValueOutOfValidRange) || (e.CharIndex != 0) {
		t.Errorf("unexpected error instance: %#v", e)
	}
	if _, err := parseElementText("13", 1, 12, 12, false); nil == err {
		t.Error("expect error")
	} else if e, ok := err.(*ErrCronExpression); !ok {
		t.Errorf("unexpected error type: %#v", err)
	} else if (e.ErrType != ErrValueOutOfValidRange) || (e.CharIndex != 1) {
		t.Errorf("unexpected error instance: %#v", e)
	}
	if _, err := parseElementText("3,13,5", 1, 12, 12, false); nil == err {
		t.Error("expect error")
	} else if e, ok := err.(*ErrCronExpression); !ok {
		t.Errorf("unexpected error type: %#v", err)
	} else if (e.ErrType != ErrValueOutOfValidRange) || (e.CharIndex != 3) {
		t.Errorf("unexpected error instance: %#v", e)
	}
}

func checkParsedElement(r *parsedElement, valuePoints uint64, lastValues uint32, weekdayValues uint32) bool {
	return (r.ValuePoints == valuePoints) && (r.LastValues == lastValues) && (r.WeekdayValues == weekdayValues)
}

func TestParseElementText02R0059(t *testing.T) {
	if r, err := parseElementText("5", 0, 59, 59, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 1<<5, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("10-15", 0, 59, 59, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 0x3f<<10, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("10-15,1", 0, 59, 59, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 0x3F<<10|1<<1, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("10-15,1,50-59/3", 0, 59, 59, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 0x3F<<10|1<<1|0x249<<50, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("10-15,1,50-59/3", 0, 59, 59, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 0x3F<<10|1<<1|0x249<<50, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
}

func TestParseElementText03R0006(t *testing.T) {
	if r, err := parseElementText("5", 0, 6, 7, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 1<<5, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("6-7", 0, 6, 7, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 0x3<<6, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("7", 0, 6, 7, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 1<<7, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("1,3L", 0, 6, 7, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 1<<1, 1<<3, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
	if r, err := parseElementText("*", 0, 6, 7, false); nil != err {
		t.Errorf("unexpected error: %v", err)
	} else if !checkParsedElement(r, 0x7F, 0, 0) {
		t.Errorf("unexpected result: %#v", r)
	}
}
