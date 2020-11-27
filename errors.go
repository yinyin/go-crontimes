package crontimes

import (
	"errors"
	"strconv"
)

//go:generate stringer -type ExpressionErrType errors.go

// ExpressionErrType defines type of errors.
type ExpressionErrType int

// Identifiers of error types.
const (
	noneExpressionErr ExpressionErrType = iota
	ErrSyntaxValueAfterLastRule
	ErrSyntaxUnexpectCharacter
	ErrSyntaxUnexpectLocation
	ErrStepMustNotBeZero
	ErrValueOutOfValidRange
	ErrLastNotSupportInMinuteRule
	ErrWeekdayNotSupportInMinuteRule
	ErrLastNotSupportInHourRule
	ErrWeekdayNotSupportInHourRule
	ErrLastDayOfMonthSyntax
	ErrLastNotSupportInMonthRule
	ErrWeekdayNotSupportInMonthRule
	ErrWeekdayNotSupportInDayOfWeekRule
)

// ErrCronExpression represent errors of cron expressions
type ErrCronExpression struct {
	ErrType   ExpressionErrType
	CharIndex int
}

func newErrCronExpression(errType ExpressionErrType, charIndex int) (e *ErrCronExpression) {
	return &ErrCronExpression{
		ErrType:   errType,
		CharIndex: charIndex,
	}
}

func (e *ErrCronExpression) Error() string {
	return "[CronExpressionError: " + e.ErrType.String() + " (" + strconv.FormatInt(int64(e.ErrType), 10) + "), Index=" + strconv.FormatInt(int64(e.CharIndex), 10) + "]"
}

// ErrBinarySizeMismatch indicate size of given binary is unexpected.
var ErrBinarySizeMismatch = errors.New("given binary size not match")
