package crontimes

type ruleType int

const (
	ruleEmpty ruleType = iota
	ruleOneValue
	ruleRange
	ruleRangeWithStep
	ruleLast
	ruleWeekday
)

type parsingState struct {
	rangeStart int
	rangeEnd   int
	stepValue  int
	ruleType   ruleType
}

func (s *parsingState) checkValue(valueMin, valueCap int) (ok bool) {
	if (s.rangeStart < valueMin) || (s.rangeStart > valueCap) {
		return false
	}
	if ((s.ruleType == ruleRange) || (s.ruleType == ruleRangeWithStep)) && ((s.rangeStart < valueMin) || (s.rangeEnd > valueCap)) {
		return false
	}
	return true
}

func parseElementText(t string, rangeMin, rangeMax, valueCap int) (result *parsedElement, err error) {
	var elem parsedElement
	var d parsingState
	t = t + ","
	for idx, ch := range t {
		switch {
		case (ch >= '0') && (ch <= '9'):
			aux := int(ch - '0')
			switch d.ruleType {
			case ruleEmpty:
				d.ruleType = ruleOneValue
				fallthrough
			case ruleOneValue:
				d.rangeStart = (d.rangeStart * 10) + aux
			case ruleRange:
				d.rangeEnd = (d.rangeEnd * 10) + aux
			case ruleRangeWithStep:
				d.stepValue = (d.stepValue * 10) + aux
			default:
				return nil, newErrCronExpression(ErrSyntaxValueAfterLastRule, idx)
			}
		case ch == '-':
			if d.ruleType == ruleOneValue {
				d.ruleType = ruleRange
			} else {
				return nil, newErrCronExpression(ErrSyntaxUnexpectCharacter, idx)
			}
		case ch == '/':
			if d.ruleType == ruleRange {
				d.ruleType = ruleRangeWithStep
			} else {
				return nil, newErrCronExpression(ErrSyntaxUnexpectCharacter, idx)
			}
		case ch == '*':
			if d.ruleType == ruleEmpty {
				d.ruleType = ruleRange
				d.rangeStart = rangeMin
				d.rangeEnd = rangeMax
			} else {
				return nil, newErrCronExpression(ErrSyntaxUnexpectCharacter, idx)
			}
		case ch == ',':
			if !d.checkValue(rangeMin, valueCap) {
				return nil, newErrCronExpression(ErrValueOutOfValidRange, idx-1)
			}
			if errType := elem.addParsingState(&d); errType != noneExpressionErr {
				return nil, newErrCronExpression(errType, idx-1)
			}
			d = parsingState{}
		case (ch == 'L') || (ch == 'l'):
			if (d.ruleType == ruleOneValue) || (d.ruleType == ruleEmpty) {
				d.ruleType = ruleLast
			} else {
				return nil, newErrCronExpression(ErrSyntaxUnexpectLocation, idx)
			}
		case (ch == 'W') || (ch == 'w'):
			if (d.ruleType == ruleOneValue) || (d.ruleType == ruleEmpty) {
				d.ruleType = ruleWeekday
			} else {
				return nil, newErrCronExpression(ErrSyntaxUnexpectLocation, idx)
			}
		case (ch == ' ') || (ch == '\t'):
		default:
			return nil, newErrCronExpression(ErrSyntaxUnexpectCharacter, idx)
		}
	}
	return &elem, nil
}

type parsedElement struct {
	ValuePoints   uint64
	LastValues    uint32
	WeekdayValues uint32
}

func (elem *parsedElement) isEmpty() bool {
	return (elem.ValuePoints == 0) && (elem.LastValues == 0) && (elem.WeekdayValues == 0)
}

func (elem *parsedElement) addParsingState(s *parsingState) (errType ExpressionErrType) {
	switch s.ruleType {
	case ruleOneValue:
		elem.ValuePoints |= (1 << s.rangeStart)
	case ruleRange:
		for i := s.rangeStart; i <= s.rangeEnd; i++ {
			elem.ValuePoints |= (1 << i)
		}
	case ruleRangeWithStep:
		if s.stepValue < 1 {
			return ErrStepMustNotBeZero
		}
		for i := s.rangeStart; i <= s.rangeEnd; i += s.stepValue {
			elem.ValuePoints |= (1 << i)
		}
	case ruleLast:
		elem.LastValues |= (1 << s.rangeStart)
	case ruleWeekday:
		elem.WeekdayValues |= (1 << s.rangeStart)
	}
	return noneExpressionErr
}
